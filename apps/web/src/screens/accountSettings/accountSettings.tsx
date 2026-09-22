import { useEffect, useRef, useState } from "react";
import StandartInput from "../../components/inputs/standartInput/standartInput";
import styles from './accountSettings.module.scss'
import { dateFormat } from "../../utils/dateFormat";
import { useNavigate } from "react-router-dom";
import MainLayout from "../../layouts/mainLayout/mainLayout";
import Icons from "../../assets/icons";
import { useDispatch, useSelector } from "react-redux";
import type { RootState } from "../../store/slices";
import ProgressBar from "../../components/progressBar/progressBar";
import { useGetProjects } from "../../hooks/useGetProjects";
import Switch from "../../components/switch/switch";
import userApi from "../../api/user.api";
import { setUser } from "../../store/slices/userSlice";
import { useMutation } from "@apollo/client/react";
import { UpdateUserDocument } from "../../graphql/generated/graphql";
import RadioButton from "../../components/radioButton/radioButton";
import StandartButton from "../../components/standartButton/standartButton";
import { setAppearance } from "../../store/slices/themeSlice";
import themeDict from "../../dicts/theme";
import buildLink from "../../utils/buildLink";

type AccentColor = "#4774fa" | "#eb5940" | "#ef6f1d" | "#715bc5" | "#4da592" | "#6bb919"

type NotificationSettingsType = {
    newComments : boolean
    issueUpdates : boolean
    mentions : boolean
    weeklyDigest : boolean
}

export type AppearanceSettingsType = {
    theme : "light" | "dark" | "system"
    accentColor : AccentColor
}

type AccountSettingsType = {
    avatar? : File
    name : string
    email : string
    login : string
}

function AccountSettingsScreen() {

    const dispatch = useDispatch()
    const [updateUser, { loading, error }] = useMutation(UpdateUserDocument)
    const user = useSelector((state : RootState) => state.user)
    const appearanceState = useSelector((state : RootState) => state.appearance)
    const [notificationSettings, setNotificationSettings] = useState<NotificationSettingsType>({
        newComments: true,
        issueUpdates: true,
        mentions: true,
        weeklyDigest: false,
    })
    const ACCENTS = ["#715bc5", "#4da592", "#4774fa", "#eb5940", "#ef6f1d", "#6bb919"] as const

    const [appearanceSettings, setAppearanceSettings] = useState<AppearanceSettingsType>(appearanceState)
    const [accountSettings, setAccountSettings] = useState<AccountSettingsType>({
        avatar : null,
        name: user.name,
        login: user.login,
        email: user.email
    })
    const fileInputRef = useRef<HTMLInputElement>(null)
    const [avatarPreviewUrl, setAvatarPreviewUrl] = useState<string>(buildLink(user.avatarKey))
    
    const notificationSettingsItems = [
        {
            id: "newComments",
            icon: <Icons.Message
                width={25}
                height={25}
                color="var(--color-text)"
            />,
            title: "New comments",
            description: "Notifications for new comments on your issues"
        },
        {
            id: "issueUpdates",
            icon: <Icons.Refresh
                width={25}
                height={25}
                color="var(--color-text)"
            />,
            title: "Status changes and assignments",
            description: "Notifications for new comments on your issues"
        },
        {
            id: "mentions",
            icon: <Icons.AtSign
                width={25}
                height={25}
                color="var(--color-text)"
            />,
            title: "Mentions",
            description: "When someone mentions you"
        },
        {
            id: "weeklyDigest",
            icon: <Icons.Mail
                width={25}
                height={25}
                color="var(--color-text)"
            />,
            title: "Weekly digest",
            description: "Summary of activity (once a week)"
        }
    ]

    const handleFileSelect = (event: React.ChangeEvent<HTMLInputElement>) => {

        const file = event.target.files?.[0]

        if(file) {
            setAccountSettings(prev => ({...prev, avatar: file}))
            
            const reader = new FileReader()

            reader.onload = (e) => {
                setAvatarPreviewUrl(e.target?.result?.toString())
            }
            reader.readAsDataURL(file)

        }

        if(fileInputRef.current) {
            fileInputRef.current.value = ""
        }

    }

    const handleAvatarClick = () => {

        fileInputRef.current?.click()

    }

    const handleSaveChanges = async () => {
        
        dispatch(setAppearance(appearanceSettings))
        
        const accountInput : {
            name? : string
            login? : string
            email? : string
        } = {}

        if(accountSettings.email != user.email) {
            accountInput.email = accountSettings.email
        }

        if(accountSettings.name != user.name) {
            accountInput.name = accountSettings.name
        }

        if(accountSettings.login != user.login) {
            accountInput.login = accountSettings.login
        }

        const hasChanges = Object.keys.length > 0

        if(hasChanges) {
            try {
                const updatedUserData = (await updateUser({
                    variables: {
                        userId: user.id,
                        input: accountInput
                    }
                })).data.updateUser

                dispatch(setUser({
                    id: user.id,
                    name: updatedUserData.name,
                    login: updatedUserData.login,
                    email: updatedUserData.email,
                    avatarKey: user.avatarKey
                }))

            }
            catch(e) {
                console.error(e)
            }
        }

        if(accountSettings.avatar) {

            try{

                const avatarKey = await userApi.setAvatar(user.id, accountSettings.avatar)

                dispatch(setUser({
                    id: user.id,
                    name: user.name,
                    login: user.login,
                    email: user.email,
                    avatarKey: avatarKey
                }))
            }
            catch(e) {
                console.error(e)
            }

        }

    }

    useEffect(() => {

        console.log(appearanceState.theme)
        
    }, [appearanceState])


    return(
        <MainLayout
            title="Settings"
            description="Manage your account, preferenses and workspace settings"
        >
            <section className={styles.main}>
                <section className={styles.accountBlock}>
                    <header>
                        <div className={styles.headerTitleBlock}>
                            <div>
                                <p className={styles.title}>Account</p>
                                <p className={styles.description}>Update your personal information and account settings</p>
                            </div>
                        </div>
                        <button className={styles.saveChangesButton} onClick={handleSaveChanges}>Save changes</button>
                    </header>
                    <div style={{display: "flex"}}>
                        <div className={styles.avatar} onClick={handleAvatarClick} style={avatarPreviewUrl == "" && !accountSettings.avatar ? {borderWidth: 1, borderColor: "var(--light-grey)", borderStyle: "solid"} : {}}>
                            {avatarPreviewUrl != ""
                                ?   <img src={avatarPreviewUrl}/>
                                :   avatarPreviewUrl != "" ? <img src={avatarPreviewUrl}/>
                                :    <Icons.PersonOutline
                                        width={30}
                                        height={30}
                                        color="var(--light-grey)"
                                    />
                            }
                            <input
                                type='file'
                                style={{display: "none"}}
                                accept="image/png,image/jpeg,image/svg+xml"
                                ref={fileInputRef}
                                onChange={handleFileSelect}
                            />
                            <div className={styles.edit}>
                                <Icons.Edit
                                    width={20}
                                    height={20}
                                    color="white"
                                />
                            </div>
                        </div>
                        <div className={styles.inputBlock}>
                            <StandartInput
                                title="Name"
                                placeholder="Name"
                                value={accountSettings.name}
                                onChange={(text) => setAccountSettings(prev => ({...prev, name: text}))}
                            />
                            <StandartInput
                                title="Email"
                                placeholder="Email"
                                value={accountSettings.email}
                                onChange={(text) => setAccountSettings(prev => ({...prev, email: text}))}
                            />
                            <StandartInput
                                title="Login"
                                placeholder="Login"
                                value={accountSettings.login}
                                onChange={(text) => setAccountSettings(prev => ({...prev, login: text}))}
                            />
                        </div>
                    </div>
                </section>
                <div className={styles.notificationAndAppearanceBlock}>
                    <section className={styles.notificationBlock}>
                        <header>
                            <div className={styles.headerTitleBlock}>
                                <Icons.BellOutline
                                    width={40}
                                    height={40}
                                    color="var(--color-accent)"
                                />
                                <div>
                                    <p className={styles.title}>Notifications</p>
                                    <p className={styles.description}>Choose what you want to be notified about</p>
                                </div>
                            </div>
                        </header>
                        {
                            notificationSettingsItems.map(setting => (
                                <div className={styles.notificationSettingsItem}>
                                    <div>
                                        {setting.icon}
                                        <div className={styles.notificationSettingsItemNamingBlock}>
                                            <p className={styles.notificationSettingsItemTitle}>
                                                {setting.title}
                                            </p>
                                            <p className={styles.notificationSettingsItemDescription}>
                                                {setting.description}
                                            </p>
                                        </div>
                                    </div>
                                    <Switch
                                        onChange={() => setNotificationSettings(prev => ({...prev, [setting.id]: !prev[setting.id]}))}
                                        enabled={notificationSettings[setting.id]}
                                    />
                                </div>
                            ))
                        }
                    </section>
                    <section className={styles.appearanceBlock}>
                        <header>
                            <div className={styles.headerTitleBlock}>
                                <Icons.Palette
                                    width={40}
                                    height={40}
                                    color="var(--color-accent)"
                                />
                                <div>
                                    <p className={styles.title}>Appearance</p>
                                    <p className={styles.description}>Customize how Buggeon looks for you</p>
                                </div>
                            </div>
                        </header>
                        <p>Theme</p>
                        <div className={styles.themeBlock}>
                            {
                                [
                                    {
                                        name: "light",
                                        icon: <Icons.Sun width={20} height={20} color="var(--light-grey)"/>
                                    },
                                    {
                                        name: "dark",
                                        icon: <Icons.MoonOutline width={20} height={20} color="var(--light-grey)"/>
                                    },
                                    {
                                        name: "system",
                                        icon: <Icons.Computer width={20} height={20} color="var(--light-grey)"/>
                                    }
                                ].map(theme => (
                                    <div className={styles.themeItem} style={{borderColor: appearanceSettings.theme == theme.name ? "var(--color-accent" : "var(--color-bg)"}} onClick={() => setAppearanceSettings(prev => ({...prev, theme: theme.name as "dark" | "light" | "system"}))}>
                                        {theme.icon}
                                        <div>
                                            <p>{theme.name}</p>
                                            <RadioButton
                                                value={appearanceSettings.theme == theme.name}
                                                onChange={() => setAppearanceSettings(prev => ({...prev, theme: theme.name as "dark" | "light" | "system"}))}
                                            />
                                        </div>
                                    </div>
                                ))
                            }
                        </div>
                        <p>Accent color</p>
                        <div style={{display: "flex", gap: 15}}>
                            {
                                ACCENTS.map(color => (
                                    <div className={styles.accentColorItem} style={{
                                        backgroundColor: color
                                    }} onClick={() => {
                                        setAppearanceSettings(prev => ({...prev, accentColor: color}))
                                    }}>
                                        {
                                            appearanceSettings.accentColor == color && <Icons.CompletedOutline width={20} height={20} color="white"/>
                                        }
                                    </div>
                                ))
                            }
                        </div>
                    </section>
                </div>
                <section className={styles.securityBlock}>
                    <header>
                        <div className={styles.headerTitleBlock}>
                            <Icons.Shield
                                width={40}
                                height={40}
                                color="var(--color-accent)"
                            />
                            <div>
                                <p className={styles.title}>Security</p>
                                <p className={styles.description}>Kepp your account safe</p>
                            </div>
                        </div>
                    </header>
                    <div className={styles.securitySettingsItem}>
                        <div className={styles.securitySettingsItemTitleBlock}>
                            <Icons.Lock                                width={25}
                                height={25}
                                color="var(--color-text)"
                            />
                            <div>
                                <p>Change password</p>
                                <p>Last changed 32 days ago</p>
                            </div>
                        </div>
                        <button className={styles.changePasswordButton}>
                            Change password
                        </button>
                    </div>
                </section>
            </section>
        </MainLayout>
    )
}

export default AccountSettingsScreen;