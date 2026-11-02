import { useNavigate, useParams } from "react-router-dom"
import { useGetMembers } from "../../hooks/useGetMembers"
import MainLayout from "../../layouts/mainLayout/mainLayout"
import styles from './diagrams.module.scss'

function DiagramsScreen() {

    return(
        <MainLayout
            title="Diagrams"
            description="Here is all project's diagrams."
        >
            <main className={styles.main}>
                
            </main>
        </MainLayout>
    )

}

export default DiagramsScreen