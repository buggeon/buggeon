import type { ProjectType } from "../../types/projectType.type";
import styles from './projectTypeLabel.module.scss'

export const colorDict : Record<ProjectType, Record<'background' | "text", string>> = {
    "Software": {
        background: "#191f3b",
        text: "#9390bf"
    },
    "Design": {
        background: "#152238",
        text: "#87bdef"
    },
    "Infrastructure": {
        background: "#13242e",
        text: "#648bb8"
    },
    "Security": {
        background: "#2c2420",
        text: "#d08a41"
    },
    "Education": {
        background: "#2a2926",
        text: "#ffdf85"
    },
    "Analytics": {
        background: "#13242e",
        text: "#648bb8"
    },
    "Integration": {
        background: "#241e30",
        text: "#b77c91"
    }
}

function ProjectTypeLabel({projectType} : {projectType : ProjectType}) {

    return(
        <div className={styles.main} style={{backgroundColor: colorDict[projectType].background}}>
            <p style={{color: colorDict[projectType].text}}>{projectType}</p>
        </div>
    )

}

export default ProjectTypeLabel