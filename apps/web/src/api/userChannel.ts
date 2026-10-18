import { WS_URL } from '../../config'

type Handler = (payload: any) => void

class UserChannel {
    private ws: WebSocket | null = null
    private handlers = new Map<string, Set<Handler>>()
    private reconnectTimer: number | null = null
    private token: string | null = null

    connect(token: string) {
        if (this.ws && this.token === token) return
        this.token = token
        this.open()
    }

    private open() {
        const url = `${WS_URL}/user?token=${encodeURIComponent(this.token!)}`
        const ws = new WebSocket(url)
        this.ws = ws

        ws.onmessage = (e) => {
            let msg: { type: string; payload: any }
            try {
                msg = JSON.parse(e.data)
            } catch {
                console.warn('[user-ws] bad payload', e.data)
                return
            }
            const set = this.handlers.get(msg.type)
            if (set) set.forEach(h => h(msg.payload))
                
        }

        ws.onclose = () => {
            this.ws = null
            this.reconnectTimer = window.setTimeout(() => this.open(), 3000)
        }

        ws.onerror = (e) => console.error('[user-ws] error', e)
    }

    disconnect() {
        if (this.reconnectTimer) clearTimeout(this.reconnectTimer)
        this.ws?.close()
        this.ws = null
        this.token = null
    }

    on(type: string, handler: Handler) {
        let set = this.handlers.get(type)
        if (!set) {
            set = new Set()
            this.handlers.set(type, set)
        }
        set.add(handler)
        return () => set!.delete(handler)
    }
}

export const userChannel = new UserChannel()