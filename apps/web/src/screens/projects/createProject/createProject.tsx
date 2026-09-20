import { useCallback, useEffect, useRef, useState } from "react"
import projectApi from "../../../api/project.api"
import DropInput from "../../../components/inputs/dropInput/dropInput"
import StandartInput from "../../../components/inputs/standartInput/standartInput"
import Modal from "../../../components/modal/modal"
import styles from './createProject.module.scss'
import type { User } from "../../../store/types/user.interface"
import { useGetProjects } from "../../../hooks/useGetProjects"
import type { RootState } from "../../../store/slices"
import { useSelector } from "react-redux"
import Icons from "../../../assets/icons"
import SystemApi from "../../../api/system.api"
import { useMutation } from "@apollo/client/react"
import { CreateProjectDocument } from "../../../graphql/generated/graphql"

interface CreateProjectProps {
    isModalVisible : boolean
    setModalVisibility
}

interface ProjectInfo {
    name : string
    description : string
    leadId : string
    logo : File
}

function CreateProject({isModalVisible, setModalVisibility} : CreateProjectProps) {

    const fileInputRef = useRef<HTMLInputElement>(null)
    const [logoPreviewUrl, setLogoPreviewUrl] = useState<string>()
    const [projectInfo, setProjectInfo] = useState<ProjectInfo | null>({
        name: "",
        description: "",
        leadId: "",
        logo: null
    })
    const [totalUsers, setTotalUsers] = useState<User[]>([])
    const user = useSelector((state : RootState) => state.user)
    const { refetch } = useGetProjects(user.id)
    const [createProject, { loading, error }] = useMutation(CreateProjectDocument)

    const handleCreateProjectButtonClick = async () => {

        try {
            const result = await createProject({
                variables: {
                    input: {
                        name: projectInfo.name,
                        description: projectInfo.description,
                        leadId: projectInfo.leadId
                    }
                }
            })

            return result.data?.createProject
        }
        catch(e) {
            console.error(`Error: ${e}`)
        }
        
    }
    

    const fetchUsers = useCallback(async () => {
        const result = await SystemApi.getAllUsers()

        if(result && result.length == 1) {
            setProjectInfo(prev => ({...prev, leadId: result[0].id}))
        }

        setTotalUsers(result)
    }, [])

    useEffect(() => {
        fetchUsers()
    }, [fetchUsers])

    const handleFileSelect = (event: React.ChangeEvent<HTMLInputElement>) => {

        const file = event.target.files?.[0]

        if(file) {
            setProjectInfo(prev => ({...prev, logo: file}))
            
            const reader = new FileReader()

            reader.onload = (e) => {
                setLogoPreviewUrl(e.target?.result?.toString())
            }
            reader.readAsDataURL(file)

        }

        if(fileInputRef.current) {
            fileInputRef.current.value = ""
        }

    }

    const handleLogoClick = () => {

        fileInputRef.current?.click()

    }

    return(
        <Modal
            onFooterActiveButtonClick={async () => {
                const project = await handleCreateProjectButtonClick()
                await projectApi.setProjectLogo(project.id, projectInfo.logo)

                setProjectInfo(null)
                setLogoPreviewUrl(null)
                await refetch()
            }}
            footerActiveButtonLabel="Create Project"
            isOpen={isModalVisible}
            setState={setModalVisibility}
            title="Create New Project"
            className={styles.createProjectModal}
        >
            <section className={styles.main}>
                <StandartInput
                    placeholder="Enter project name..."
                    subtitle='Choose a clear, descriptive name for your project'
                    title="Project Name"
                    onChange={(text) => {setProjectInfo(prev => ({...prev, name: text}))}}
                />
                <StandartInput
                    placeholder="Describe your project, its goals, and key objectives..."
                    subtitle='Provide a brief description of what this project is about'
                    title="Description"
                    onChange={(text) => {setProjectInfo(prev => ({...prev, description: text}))}}
                    style={{height: 100}}
                />
                {
                    projectInfo?.logo
                        ?   <div className={styles.logoPicker} onClick={handleLogoClick}>
                                <img src={logoPreviewUrl}/>
                                <input 
                                    type='file'
                                    style={{display: "none"}}
                                    accept="image/png,image/jpeg,image/svg+xml"
                                    ref={fileInputRef}
                                    onChange={handleFileSelect}
                                />
                            </div>
                        :   <div className={styles.logoPicker} onClick={handleLogoClick}>
                                <input 
                                    type='file'
                                    style={{display: "none"}}
                                    accept="image/png,image/jpeg,image/svg+xml"
                                    ref={fileInputRef}
                                    onChange={handleFileSelect}
                                />
                                <Icons.Upload color='var(--dark-grey)' width={30} height={30}/>
                                <p>Upload logo</p>
                                <p>PNG, JPG or SVG</p>
                                <p>Max size 2MB</p>
                            </div>
                }
                <DropInput
                    value={`${totalUsers.filter(user => user.id == projectInfo?.leadId)[0]?.name}${totalUsers.length == 1 && projectInfo?.leadId == user.id ? " (You)" : ""}`}
                    placeholder="Select project lead..."
                    subtitle='Choose who will be leading this project'
                    onChange={(user) => {setProjectInfo(prev => ({...prev, leadId: user.id}))}}
                    title="Project Lead"
                    content={totalUsers.map(user => ({
                        name: user.name,
                        id: user.id
                    }))}
                    renderFunc={(item) => (
                        <div>
                            <p>{item.name}</p>
                        </div>
                    )}
                />
            </section>
        </Modal>
    )
}

export default CreateProject