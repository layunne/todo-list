import { Task } from "./task";

export interface Project {
  id: string;
  name: string;
  userId: string;
  tasks: Task[];
}

export interface CreateProject {
  name: string;
}

export interface UpdateProject {
  id: string;
  name: string;
}





