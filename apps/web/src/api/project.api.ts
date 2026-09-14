import type { ProjectModel } from "../models/project.model";
import type { SchemaModel } from "../models/schema.model";
import api from "./api";
import type { CardData } from "./types/createCard.type";

class ProjectApi {
    
    async setProjectLogo(projectId : string, logo : File) {

        const formData = new FormData()
        formData.append("logo", logo)

        try {
            await api.patch(`/api/projects/${projectId}/logo`, formData, {
                headers: {
                    "Content-Type": "multipart/form-data"
                }
            })
        }
        catch(e) {

        }

    }

    async updateCardLocation(projectId : string, cardId : string, oldBoardId : string, newBoardId : string) {

        try{

            const result = await api.put(`/api/projects/${projectId}/boards/${oldBoardId}/cards/${cardId}/updatelocation`, { newBoardId })

            if(result.status != 200) {
                throw new Error("Failed to update card location")
            }

        }
        catch(e) {
            throw new Error(e)
        }
        
    }

}

const projectApi = new ProjectApi()

export default projectApi