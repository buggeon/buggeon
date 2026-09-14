import styles from './entry.module.scss'
import { useState } from 'react';
import EntryForm from './entryForm/entryForm';
import Icons from '../../assets/icons';

function EntryScreen() {

    const [mode, setMode] = useState<"auth" | "regist">("auth")

    const onboardingTags = [
        {
            icon: <Icons.Checkbox width={40} height={40} color='#8a59f7'/>,
            title: "Organize Issues",
            description: "Keep everyting in one place and never miss a critical bug."
        },
        {
            icon: <Icons.TeamOutline width={40} height={40} color='#8a59f7'/>,
            title: "Collaborate Seamlessly",
            description: "Work together with your team in real time and ship faster."
        },
        {
            icon: <Icons.StatisticUp width={40} height={40} color='#8a59f7'/>,
            title: "Gain Insights",
            description: "Visualize progess make data-driven decisions."
        }
    ]

    return(
        <section className={styles.main}>
            <section className={styles.onboardingBlock}>
                <div className={styles.logo}>
                    <Icons.Logo width={50} height={50}/>
                    <h1>Buggeon</h1>
                </div>
                <div className={styles.wellcomeLabel}>
                    <p>Welcome to Buggeon</p>
                </div>
                <div className={styles.title}>
                    <p>
                        Track. Prioritize.{"\t"}
                        <span>
                            Resolve.{"\t"}
                        </span>
                        Ship better software
                    </p>
                </div>
                <div className={styles.descriptionBlock}>
                    <p>
                        Buggeon helps teams manage issues, collaborate seamlessly, and deliver quality product faster.
                    </p>
                </div>
                <div className={styles.tagsBlock}>
                    {
                        onboardingTags.map(tag => (
                            <div className={styles.tag}>
                                <div className={styles.tagIconBlock}>
                                    {tag.icon}
                                </div>
                                <div>
                                    <p className={styles.tagTitle}>{tag.title}</p>
                                    <p className={styles.tagDescription}>{tag.description}</p>
                                </div>
                            </div>
                        ))
                    }
                </div>
                <div className={styles.securityBlock}>
                    <div className={styles.lockBlock}>
                        <Icons.Lock2 width={20} height={20} color='#b7becc'/>
                    </div>
                    <p className={styles.aboutSecurity}>
                        Your data is secure with enterprise-grade security and privacy
                    </p>
                </div>
            </section>
            <EntryForm key={mode} mode={mode} changeMode={setMode} className={styles.entryFormBlock}/>
        </section>
    )
}

export default EntryScreen;