import { useCallback, useEffect, useRef, useState } from 'react'

export type ChatMessage = {
    id?: string
    content: string
    senderId?: string
    senderName?: string
    senderAvatarUrl?: string
    sentAt?: string
}

type ReadyState = 'connecting' | 'open' | 'closing' | 'closed'

type UseCardChatOptions = {
    cardId: string
    onMessage?: (msg: ChatMessage) => void
}

export function useCardChat({ cardId, onMessage }: UseCardChatOptions) {
    const [readyState, setReadyState] = useState<ReadyState>('closed')
    const [lastMessage, setLastMessage] = useState<ChatMessage | null>(null)

    const wsRef = useRef<WebSocket | null>(null)
    const onMessageRef = useRef(onMessage)
    useEffect(() => { onMessageRef.current = onMessage }, [onMessage])

    useEffect(() => {
        const token = localStorage.getItem('accessToken')
        if (!token) {
            return
        }

        const apiUrl = import.meta.env.WS_URL as string
        const url = `${apiUrl}/${cardId}?token=${encodeURIComponent(token)}`

        const ws = new WebSocket(url)
        wsRef.current = ws

        ws.onopen = () => {
            console.log('[ws] connected to', cardId)
            setReadyState('open')
        }

        ws.onmessage = (event) => {
            let parsed = JSON.parse(event.data)
            setLastMessage(parsed)
            onMessageRef.current?.(parsed)
        }

        ws.onerror = (e) => {
            console.error('[ws] error', e)
        }

        ws.onclose = (e) => {
            console.log('[ws] closed', e.code, e.reason)
            setReadyState('closed')
        }

        return () => {
            ws.onopen = null
            ws.onmessage = null
            ws.onerror = null
            ws.onclose = null
            ws.close()
            wsRef.current = null
        }
    }, [cardId])

    const send = useCallback((data: ChatMessage | string) => {
        const ws = wsRef.current
        if (!ws || ws.readyState !== WebSocket.OPEN) {
            console.warn('[ws] cannot send, connection not open')
            return false
        }
        const payload = typeof data === 'string' ? data : JSON.stringify(data)
        ws.send(payload)
        return true
    }, [])

    return {
        lastMessage,
        readyState,
        connected: readyState === 'open',
        send,
    }
}