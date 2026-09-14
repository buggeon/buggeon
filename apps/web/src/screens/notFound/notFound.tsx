import Icons from '../../assets/icons'
import styles from './notFound.module.scss'

function NotFoundScreen() {
    return(
        <section className={styles.main}>
            <Icons.Logo width={50} height={50}/>
            <p>Oops... Looks like you took the wrong path</p>
        </section>
    )   
}

export default NotFoundScreen