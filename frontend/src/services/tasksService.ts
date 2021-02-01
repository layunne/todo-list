import { CreateTask, UpdateTask, ChangeStatusTask } from "../models/task";
import api from "./api";

class TasksService {

  static get =  async (taskId: String) =>  {
    return api.get(`/tasks/${taskId}`)
  }

  static getAll =  async () =>  {
    return api.get(`/tasks`)
  }

  static create =  async (task: CreateTask) =>  {
    return api.post(`/tasks`, task)
  }

  static update =  async (task: UpdateTask) =>  {
    return api.put(`/tasks`, task)
  }

  static changeStatus =  async (status: ChangeStatusTask) =>  {
    return api.put(`/tasks/status`, status)
  }

  static delete =  async (taskId: String) =>  {
    return api.delete(`/tasks/${taskId}`)
  }
}


export default TasksService;