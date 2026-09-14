import { useCallback, useEffect, useRef, useState } from "react";
import StandartInput from "../../components/inputs/standartInput/standartInput";
import styles from './projects.module.scss'
import type { ProjectModel } from "../../models/project.model";
import { dateFormat } from "../../utils/dateFormat";
import ProjectTypeLabel, { colorDict } from "../../components/projectTypeLabel/projectTypeLabel";
import { useNavigate } from "react-router-dom";
import MainLayout from "../../layouts/mainLayout/mainLayout";
import Icons from "../../assets/icons";
import projectApi from "../../api/project.api";
import { useSelector } from "react-redux";
import type { RootState } from "../../store/slices";
import ProgressBar from "../../components/progressBar/progressBar";
import { useGetProjects } from "../../hooks/useGetProjects";
import Modal from "../../components/modal/modal";
import DropInput from "../../components/inputs/dropInput/dropInput";
import SystemApi from "../../api/system.api";
import type { User } from "../../store/types/user.interface";
import CreateProject from "./createProject/createProject";

interface ProjectInfo {
    name : string
    description : string
    leadId : string
    logo : File
}

function ProjectsScreen() {

    const [findName, setFindName] = useState("")
    const [page, setPage] = useState(1)
    const navigate = useNavigate()
    const user = useSelector((state : RootState) => state.user)

    const { data, refetch } = useGetProjects(user.id)
    const [isModalVisible, setModalVisibility] = useState(false)

    return(
        <MainLayout
            title="Projects"
            description="View and manage all your project in one place"
        >
            <section className={styles.main}>
                <CreateProject isModalVisible={isModalVisible} setModalVisibility={setModalVisibility}/>
                <section className={styles.header}>
                    <button className={styles.newProjectButton} onClick={() => {
                        setModalVisibility(true)
                    }}>
                        <Icons.Plus width={20} height={20} color="#FFFFFF" className={styles.newProjectButtonPlus}/>
                        <p>New Project</p>
                    </button>
                    <div>
                        <StandartInput style={{height: 50}} prefixIcon={<Icons.Search width={20} height={20} color="#b7becc"/>} placeholder="Search projects..." onChange={text => setFindName(text)}/>
                        <button className={styles.filterAndSortButton}>
                            <div>
                                <Icons.FilterOutline width={20} height={20} color="#b7becc"/>
                            </div>
                            <p>Filter</p>
                        </button>
                        <button className={styles.filterAndSortButton}>
                            <div>
                                <Icons.Sort width={20} height={20} color="#b7becc"/>
                            </div>
                            <p>Sort</p>
                        </button>
                        <div className={styles.visualTypeBlock}>
                            <button>
                                <Icons.Widgets width={20} height={20} color="#b7becc"/>
                            </button>
                            <button>
                                <Icons.List width={20} height={20} color="#b7becc"/>
                            </button>
                        </div>
                    </div>
                </section>
                {
                    data && data.projects.length != 0
                    ?
                    <section>
                        <section className={styles.listBlock}>
                            <table>
                                <thead>
                                    <tr>
                                        <th>Name</th>
                                        <th>Lead</th>
                                        <th>Members</th>
                                        <th>Issues</th>
                                        <th>Progress</th>
                                        <th>Updated</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {
                                        data.projects.filter(project => project.name.toLowerCase().includes(findName.toLowerCase()) || project.description.toLowerCase().includes(findName.toLowerCase())).map(project => (
                                            <tr onClick={() => {
                                                navigate(`/project/${project.id}`)
                                            }} key={project.id}>
                                                <td>
                                                    <div className={styles.projectAbout}>
                                                        <img src={project.logoUrl}/>
                                                        <div className={styles.projectNameAndDescription}>
                                                            <p>{project.name}</p>
                                                            <p>{project.description}</p>
                                                        </div>
                                                    </div>
                                                </td>
                                                <td>
                                                    <div className={styles.leadInfo}>
                                                        <img src={project.logoUrl}/>
                                                        <p>{project.lead.user.name}</p>
                                                    </div>
                                                </td>
                                                <td>
                                                    <div className={styles.members}>
                                                        <img src=""/>
                                                        <img src=""/>
                                                        <img src=""/>
                                                        <div>
                                                            <p>+5</p>
                                                        </div>
                                                    </div>
                                                </td>
                                                <td>
                                                    <div className={styles.issues}>
                                                        <p>0</p>
                                                        <p>Open</p>
                                                    </div>
                                                </td>
                                                <td>
                                                    <p>{project.progress}%</p>
                                                    <ProgressBar progress={project.progress} color="var(--foreground-blue)" className={styles.progressBar}/>
                                                </td>
                                                <td>
                                                    <p className={styles.updatedAt}>{dateFormat(Date.parse(project.updatedAt))}</p>
                                                </td>
                                            </tr>
                                        ))
                                    }
                                </tbody>
                            </table>
                        </section>
                        <section className={styles.pagSwitch}>
                            <button>
                                <Icons.Arrow width={20} height={20} color="#b7becc"/>
                            </button>
                            <div className={styles.pagNumber}>
                                <p>{page}</p>
                            </div>
                            <button>
                                <Icons.Arrow width={20} height={20} color="#b7becc"/>
                            </button>
                        </section>
                    </section>
                    : <p className={styles.noProjectsYet}>No projects yet...</p>
                }
            </section>
        </MainLayout>
    )
}

export default ProjectsScreen;