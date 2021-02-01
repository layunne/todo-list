import { CreateProject, UpdateProject } from "../models/project";
import api from "./api";

class ProjectsService {
  
  static get =  async (projectId: String) =>  {
    return api.get(`/projects/${projectId}`)
  }

  static getAll =  async () =>  {
    return api.get(`/projects`)
  } 

  static create =  async (project: CreateProject) =>  {
    return api.post(`/projects`, project)
  }

  static update =  async (project: UpdateProject) =>  {
    return api.put(`/projects`, project)
  }

  static delete =  async (projectId: String) =>  {
    return api.delete(`/projects/${projectId}`)
  }

}


export default ProjectsService;