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


function ProfileScreen() {

    const user = useSelector((state : RootState) => state.user)
    const { data } = useGetProjects(user.id)

    return(
        <MainLayout
            title={`Welcome back, ${user.name}`}
            description="Here's what's happening with your projects today."
        >
            <section className={styles.main}>
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