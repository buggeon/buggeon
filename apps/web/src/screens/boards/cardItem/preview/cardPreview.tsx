import Mark from '../../../../components/mark/mark';
import { priorityDict } from '../../../../dicts/priority';
import styles from './cardPreview.module.scss'
import { useNavigate } from 'react-router-dom';

function CardPreview({title, priority, status, themeColor, boardId, cardId} : {title : string, priority : string, status : string, themeColor : string, boardId : string, cardId : string}) {

    const navigate = useNavigate()

    return(
        <section className={styles.main} style={{borderLeftColor: themeColor}} onClick={() => navigate(`${boardId}/cards/${cardId}`)}>
            <p>{title}</p>
            {
                priority != "None" && <Mark backgroundColor={status == "Done" ? priorityDict["Done"].backgroundColor : priorityDict[priority].backgroundColor} foregroundColor={status == "Done" ? priorityDict["Done"].foregroundColor : priorityDict[priority].foregroundColor} title={priority}/>
            }
        </section>
    )
}

export default CardPreview;