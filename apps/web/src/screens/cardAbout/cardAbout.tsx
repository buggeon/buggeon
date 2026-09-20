import styles from './cardAbout.module.scss'
import MainLayout from '../../layouts/mainLayout/mainLayout';
import { useParams } from 'react-router-dom';
import { useGetCard } from '../../hooks/useGetCard';
import Icons from '../../assets/icons';
import { useState } from 'react';
import { useCardChat } from '../../hooks/useCardChat';
import type { ChatMessage } from '../../hooks/useCardChat';
import { useSelector } from 'react-redux';
import type { RootState } from '../../store/slices';

function CardAboutScreen() {

    const {cardId} = useParams()
    const user = useSelector((state : RootState) => state.user)
    const { data } = useGetCard(cardId, ["title", "content", "messages { content, sender { id , name, avatarUrl } }"])
    const [messageText, setMessageText] = useState("")
    const [liveMessages, setLiveMessages] = useState<ChatMessage[]>([])

    const history: ChatMessage[] = data?.card.messages.map(m => ({
        senderId: m.sender.id,
        senderName: m.sender.name,
        senderAvatarUrl: m.sender.avatarUrl,
        content: m.content,
        id: m.id,
    })) ?? []

    const { send } = useCardChat({
        cardId,
        onMessage: (msg) => {
            setLiveMessages(prev => [...prev, msg])
        }
    })

    const messages = [...history, ...liveMessages]

    return(
        <MainLayout
            title={data?.card?.title}
            description=""
            isModal
        >
            <section className={styles.main}>
                <section>
                    <p>{data?.card?.content}</p>
                </section>
                <section className={styles.discussion}>
                    <header>
                        <Icons.Message width={30} height={30} color='white'/>
                        <p>Discussion</p>
                    </header>
                    <section className={styles.chatBlock}>
                        {
                            messages && messages?.map(message => (
                                <div className={styles.messageItem} style={{justifyContent: user.id == message.senderId ? "flex-end" : "flex-start"}}>
                                    <p>{message.content}</p>
                                </div>
                            ))
                        }
                        <div className={styles.messageInput}>
                            <Icons.Smile width={25} height={25} color='white'/>
                            <textarea onChange={(event) => setMessageText(event.target.value)}/>
                            <Icons.Send onClick={() => {
                                send(messageText)
                                setMessageText("")
                            }} width={25} height={25} color='white'/>
                        </div>
                    </section>
                </section>
            </section>
        </MainLayout>
    )
}

export default CardAboutScreen;