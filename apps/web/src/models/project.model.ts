import type { ProjectType } from "../types/projectType.type"

export interface ProjectModel {
    id : string
    name : string
    description : string
    type : ProjectType
    leadId : string
    membersIds : string[]
    issues : number
    progress : number
    updatedAt : number
    logoUrl : string 
}