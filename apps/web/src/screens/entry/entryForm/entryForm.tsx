import { useState, type Dispatch, type SetStateAction } from 'react'
import styles from './entryForm.module.scss'
import StandartInput from '../../../components/inputs/standartInput/standartInput'
import userApi from '../../../api/user.api'
import { useNavigate } from 'react-router-dom'
import Icons from '../../../assets/icons'
import { useDispatch } from 'react-redux'
import { setUser } from '../../../store/slices/userSlice'
import type { User } from '../../../store/types/user.interface'

function EntryForm({mode, changeMode, className} : {mode : "regist" | "auth", changeMode : Dispatch<SetStateAction<"auth" | "regist">>, className? : string}) {

    const [isPasswordVisible, setPasswordVisibility] = useState(false)
    const [login, setLogin] = useState("@")
    const [password, setPassword] = useState("")
    const [loading, setLoading] = useState(false)
    const [name, setName] = useState("")
    const [email, setEmail] = useState("")
    const dispatch = useDispatch()

    const navigate = useNavigate()
    
    const alternativeEntryVariants = [
        {
            icon: <Icons.Google width={30} height={30}/>,
            title: "Google",
            onClick: () => {}
        },
        {
            icon: <Icons.GitHub width={30} height={30}/>,
            title: "GitHub",
            onClick: () => {}
        }
    ]

    const handleLoginChange = (text: string) => {
        let value = text;
        
        if (!value.startsWith('@')) {
            value = '@' + value.replace(/@/g, '');
        } else {
            const atIndex = value.indexOf('@');
            if (atIndex === 0) {
                value = '@' + value.slice(1).replace(/@/g, '');
            } else {
                value = '@' + value.replace(/@/g, '');
            }
        }
        
        setLogin(value);
    };

    return(
        <section className={`${styles.authForm} ${className}`}>
            <h1 className={styles.authFormTitle}>
                {mode == "auth" ? "Welcome back" : "Create your account"}
            </h1>
            <p className={styles.authFormDescription}>
                {mode == "auth" ? "Log in to your Buggeon account" : "Join your team and start tracking issues"}
            </p>
            <div className={styles.alternativeEntryVarinatsButtonBlock}>
                {
                    alternativeEntryVariants.map(variant =>(
                        <button onClick={variant.onClick} className={styles.alternativeEntryVariantButton}>
                            {variant.icon}
                            <p>Continue with {variant.title}</p>
                        </button>
                    ))
                }
            </div>
            <div className={styles.alternativeBlock}>
                <hr></hr>
                <p>or</p>
                <hr></hr>
            </div>
            {
                mode == "auth" 
                    ?   <div className={styles.inputBlock}>
                            <StandartInput isNecessary title='Login' prefixIcon={<Icons.Mail width={20} height={20} color='#FFFFFF'/>} placeholder='Enter your login ( starts with "@" )' onChange={(text => handleLoginChange(text))} value={login}/>
                            <StandartInput isNecessary title='Password' prefixIcon={<Icons.Lock width={20} height={20} color='#FFFFFF'/>} suffixIcon={isPasswordVisible ? <Icons.EyeOff width={20} height={20} color='#FFFFFF'/> : <Icons.Eye width={20} height={20} color='#FFFFFF'/>} placeholder='Enter your password' onChange={(text => setPassword(text))} onSuffixClick={() => {
                                setPasswordVisibility(!isPasswordVisible)
                            }} type={!isPasswordVisible && 'password'}/>
                        </div>
                    :   <div className={styles.inputBlock}>
                            <StandartInput title='Name' prefixIcon={<Icons.PersonOutline width={20} height={20} color="#FFFFFF"/>} placeholder='Enter your name' onChange={(text) => setName(text)}/>
                            <StandartInput title='Login' prefixIcon={<Icons.Mail width={20} height={20} color="#FFFFFF"/>} placeholder='Invite a login' onChange={(text) => setLogin(text)}/>
                            <StandartInput title='Email' prefixIcon={<Icons.PersonOutline width={20} height={20} color="#FFFFFF"/>} placeholder='Enter your contact email' onChange={(text) => setEmail(text)} value={email}/>
                            <StandartInput title='Password' prefixIcon={<Icons.Lock width={20} height={20} color="#FFFFFF"/>} suffixIcon={isPasswordVisible ? <Icons.EyeOff width={20} height={20} color='#FFFFFF'/> : <Icons.Eye width={20} height={20} color='#FFFFFF'/>} placeholder='Create a strong password' onSuffixClick={() => {
                                setPasswordVisibility(!isPasswordVisible)
                            }} type={!isPasswordVisible && 'password'} onChange={(text) => setPassword(text)}/>
                        </div>
            }
            <button className={styles.entryButton} onClick={async () => {
                
                if(mode == "auth") {
                    if(login != "" && password != "") {
                        setLoading(true)
                        try{
                            const result = await userApi.login(login, password)
                        
                            dispatch(setUser({
                                id: result.id,
                                name: result.name,
                                login: result.login,
                                email: result.email,
                                avatarKey: result.avatarKey
                            }))
                            setLoading(false)
                            navigate("/dashboard/profile")
                        }
                        catch(e) {
                            setLoading(false)
                        }
                    }
                }
                else {
                    if(login != "" && name != "" || password != "" || email != "") {
                        try{
                            setLoading(true)
                            const result = await userApi.regist(login, password, email, name)
                            dispatch(setUser({
                                id: result.id,
                                name: result.name,
                                login: result.login,
                                email: result.email,
                                avatarKey: result.avatarKey
                            } as User))
                            setLoading(false)
                            navigate("/dashboard/profile")
                        }
                        catch(e) {
                            setLoading(false)
                        }
                    }
                }
                
            }}>
                {loading ? <Icons.Spinner width={30} height={30} className={styles.spinner}/> : mode == "auth" ? "Sign in" : "Sign up"}
            </button>
            <p className={styles.signin}>
                {mode == "auth" ? "Don't have an account?" : "Already have an account?"}{"\t"}
                <span onClick={() => {
                    changeMode(mode == "auth" ? "regist" : "auth")
                }}>
                    {mode == "auth" ? "Sign up" : "Sign in"}
                </span>
            </p>
        </section>
    )
}

export default EntryForm;