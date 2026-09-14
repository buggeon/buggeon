import type { IInfoCard } from './infoCard/infoCard';
import InfoCard from './infoCard/infoCard';
import styles from './profile.module.scss'
import MainLayout from '../../layouts/mainLayout/mainLayout';
import Icons from '../../assets/icons';
import { useSelector } from 'react-redux';
import type { RootState } from '../../store/slices';
import { Line, LineChart, ResponsiveContainer, XAxis, YAxis } from 'recharts';
import ProgressBar from '../../components/progressBar/progressBar';
import { useGetProjects } from '../../hooks/useGetProjects';

function Icon({icon, backgroundColor} : {icon : React.ReactNode, backgroundColor : string}) {
    return(
        <div style={{backgroundColor: backgroundColor}} className={styles.infoCardIcon}>
            {icon}
        </div>
    )
}

function ProfileScreen() {

    const user = useSelector((state : RootState) => state.user)
    const { data } = useGetProjects(user.id)

    const infoCards : IInfoCard[] = [
        {
            title: "My Projects",
            value: "7",
            additionalInfo: "+2 from last month",
            additionalInfoTextColor: "var(--foreground-blue)",
            icon: Icon({
                icon: <Icons.FoldersFilled width={30} height={30} color="var(--foreground-blue)"/>,
                backgroundColor: "var(--background-blue)"
            })
        },
        {
            title: "Open Issues",
            value: "3",
            additionalInfo: "+1 from last week",
            additionalInfoTextColor: "var(--foreground-red)",
            icon: Icon({
                icon: <Icons.BugFilled width={30} height={30} color="var(--foreground-red)"/>,
                backgroundColor: "var(--background-red)"
            })
        },
        {
            title: "Total Tasks",
            value: "12",
            additionalInfo: "+3 from this month",
            additionalInfoTextColor: "var(--foreground-purple)",
            icon: Icon({
                icon: <Icons.TaskFilled width={30} height={30} color="var(--foreground-purple)"/>,
                backgroundColor: "var(--background-purple)"
            })
        },
        {
            title: "Completed Tasks",
            value: "72%",
            additionalInfo: "+8 from last week",
            additionalInfoTextColor: "var(--foreground-green)",
            icon: Icon({
                icon: <Icons.CompletedTaskFilled width={30} height={30} color="var(--foreground-green)"/>,
                backgroundColor: "#183033"
            })
        },
    ]

    return(
        <MainLayout
            title={`Welcome back, ${user.name}`}
            description="Here's what's happening with your projects today."
        >
            <section className={styles.main}>
                <section className={styles.infoCardBlock}>
                    {
                        infoCards.map(card => (
                            <InfoCard {...card}/>
                        ))
                    }
                </section>
                <section className={styles.statisticAndProgress}>
                    <section className={styles.card}>
                        <header>
                            <h4>Activity Overview</h4>
                        </header>
                        <ResponsiveContainer>
                            <LineChart data={[
                                {name: "Янв.", openedIssues: 40, closedIssues: 20},
                                {name: "Февр.", openedIssues: 45, closedIssues: 30},
                                {name: "Мар.", openedIssues: 60, closedIssues: 50},
                                {name: "Апр.", openedIssues: 10, closedIssues: 0},
                                {name: "Май.", openedIssues: 30, closedIssues: 10},
                            ]}>
                                <YAxis tick={{
                                    fill: "#FFFFFF"
                                }} axisLine={{
                                    stroke: "var(--dark-grey)",
                                    strokeWidth: 0.5
                                }}/>
                                <XAxis dataKey="name" tick={{
                                    fill: "#FFFFFF"
                                }} axisLine={{
                                    stroke: "var(--dark-grey)",
                                    strokeWidth: 0.5
                                }}/>
                                <Line dataKey="openedIssues" stroke="var(--foreground-blue)"/>
                                <Line dataKey="closedIssues" stroke="var(--foreground-green)"/>
                            </LineChart>
                        </ResponsiveContainer>
                    </section>
                    <section className={styles.card}>
                        <header>
                            <h4>Projects status</h4>
                            <p>View all</p>
                        </header>
                        <div className={styles.projectsBlock}>
                            {
                                data && data.projects.map(project => (
                                    <div className={styles.projectItem}>
                                        <div className={styles.nameAndLogoBlock}>
                                            <img src={project.logoUrl}/>
                                            <p>{project.name}</p>
                                        </div>
                                        <div className={styles.progressBarBlock}>
                                            <ProgressBar progress={project.progress} color='var(--foreground-purple)' className={styles.progressBar}/>
                                            <p>{project.progress}%</p>
                                        </div>
                                    </div>
                                ))
                            }
                        </div>
                    </section>
                </section>
                <section className={styles.activityAndTasks}>
                    <section className={styles.card}>
                        <header>
                            <h4>Recent Activity</h4>
                            <p>View all</p>
                        </header>
                    </section>
                    <section className={styles.card}>
                        <header>
                            <h4>My Tasks</h4>
                            <p>View all</p>
                        </header>
                    </section>
                </section>
            </section>
        </MainLayout>
    )
}

export default ProfileScreen;