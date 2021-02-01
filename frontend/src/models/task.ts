
export interface Task {
  id: string;
  projectId: string;
  description: string;
  status: boolean;
  createdAt: number;
  finishedAt: number;
  toFinishAt: number;
}

export interface CreateTask {
  projectId: string;
  description: string;
  toFinishAt: number;
}

export interface UpdateTask {
  id: string;
  description: string;
  toFinishAt: number;
}

export interface ChangeStatusTask {
  id: string;
  status: boolean;
}
