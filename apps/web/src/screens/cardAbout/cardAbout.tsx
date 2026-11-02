import styles from './cardAbout.module.scss'
import MainLayout from '../../layouts/mainLayout/mainLayout';
import { useParams } from 'react-router-dom';
import { useGetCard } from '../../hooks/useGetCard';
import Icons from '../../assets/icons';
import { useState } from 'react';
import { useCardChannel } from '../../hooks/useCardChannel';
import type { ChatMessage } from '../../hooks/useCardChannel';
import { useSelector } from 'react-redux';
import type { RootState } from '../../store/slices';
import buildLink from '../../utils/buildLink';
import AvatarImage from '../../components/avatarImage/avatarImage';

function CardAboutScreen() {

    const {cardId} = useParams()
    const user = useSelector((state : RootState) => state.user)
    const { cardData, loading } = useGetCard(cardId, ["title", "content", "messages { content, sender { id , name, avatarKey } }"])
    const [messageText, setMessageText] = useState("")
    const [liveMessages, setLiveMessages] = useState<ChatMessage[]>([])

    const history: ChatMessage[] = cardData?.card.messages.map(m => ({
        senderId: m.sender.id,
        senderName: m.sender.name,
        senderAvatarKey: m.sender.avatarKey,
        content: m.content,
        id: m.id,
    })) ?? []

    const { send } = useCardChannel({
        cardId,
        onMessage: (msg) => {
            setLiveMessages(prev => [...prev, msg])
        }
    })

    const messages = [...history, ...liveMessages]

    return(
        <MainLayout
            title={cardData?.card?.title}
            description=""
            isModal
            loading={loading}
        >
            <section className={styles.main}>
                <section>
                    <p>{cardData?.card?.content}</p>
                </section>
                <section className={styles.discussion}>
                    <header>
                        <Icons.Message width={30} height={30} color='var(--color-text)'/>
                        <p>Discussion</p>
                    </header>
                    <ul className={styles.chatBlock}>
                        {
                            messages && messages?.map(message => (
                                <li className={styles.messageItem} style={{justifyContent: user.id == message.senderId ? "flex-end" : "flex-start"}}>
                                    <p style={{order: user.id == message.senderId ? -1 :  1}}>{message.content}</p>
                                    <AvatarImage
                                        imageKey={message.senderAvatarKey}
                                        alt={message.senderName}
                                        size={30}
                                    />
                                </li>
                            ))
                        }
                    </ul>
                    <div className={styles.messageInputBlock}>
                        <Icons.Smile width={25} height={25} color='var(--color-text)'/>
                        <textarea
                            onChange={(event) => setMessageText(event.target.value)}
                            value={messageText}
                        />
                        <Icons.Send onClick={() => {
                            send(messageText)
                            setMessageText("")
                        }} width={25} height={25} color='var(--color-text)'/>
                    </div>
                </section>
            </section>
        </MainLayout>
    )
}

export default CardAboutScreen;