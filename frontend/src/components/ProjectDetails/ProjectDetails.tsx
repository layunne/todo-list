import React from 'react';

import { createStyles, Theme, makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import Typography from '@material-ui/core/Typography';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import { Project } from '../../models/project';
import { Task, UpdateTask } from '../../models/task';
import TaskItem from '../TaskItem';
import TasksService from '../../services/tasksService';
import ProjectsService from '../../services/projectsService';
import AddTask from '../AddTask';
import IconButton from '@material-ui/core/IconButton';
import DeleteIcon from '@material-ui/icons/Delete';
import EditeIcon from '@material-ui/icons/EditRounded';
import AddEditProject from '../AddEditProject';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    toolbar: theme.mixins.toolbar,
    content: {
      flexGrow: 1,
      backgroundColor: theme.palette.background.default,
      padding: theme.spacing(3),
    },
    table: {
      minWidth: 650,
    },
  }),
);

type ProjectProps = {
  project: Project;
  handlerUpdateProject: (project: Project | undefined) => void;
};

const ProjectDetails: React.FC<ProjectProps> = ({ project, handlerUpdateProject }: ProjectProps) => {
  const classes = useStyles();

  const [projectState, setProjectState] = React.useState<Project>(project);
  const [isEditProject, setIsEditProject] = React.useState(false);

  const updateProject = async () => {
    const response = await ProjectsService.get(project.id);
    project = response.data;

    setProjectState(response.data);
    handlerUpdateProject(response.data);
  };

  const handleStatusChange = async (id: string, status: boolean) => {
    await TasksService.changeStatus({ id, status });

    await updateProject();
  };

  const handlerDelete = async (taskId: string) => {
    await TasksService.delete(taskId);

    const response = await ProjectsService.get(project.id);

    handlerUpdateProject(response.data);
  };

  const handlerEditTask = async (task: UpdateTask) => {
    await TasksService.update(task);

    const response = await ProjectsService.get(project.id);

    handlerUpdateProject(response.data);
  };

  const onChangeInputAdd = async () => {
    const response = await ProjectsService.get(project.id);

    handlerUpdateProject(response.data);
  };

  const onDeleteProject = async () => {
    await ProjectsService.delete(project.id);

    handlerUpdateProject(undefined);
  };
  const handleEditeProject = (project: Project | undefined) => {
    setIsEditProject(false);
    if (project) {
      handlerUpdateProject(project);
    }
  };

  const todoList = project.tasks.filter((t: Task) => !t.status);
  const doneList = project.tasks.filter((t: Task) => t.status);

  return (
    <main className={classes.content}>
      <div className={classes.toolbar} />

      {isEditProject && <AddEditProject project={project} handlerUpdateProject={handleEditeProject} />}

      <div style={{ display: 'flex', margin: '10px' }}>
        <Typography
          variant="h6"
          color="inherit"
          noWrap
          style={{
            flexGrow: 1,
            margin: '30px',
          }}
        >
          {project.name}
        </Typography>
        <IconButton color="inherit" onClick={() => setIsEditProject(true)}>
          <EditeIcon />
        </IconButton>{' '}
        <IconButton onClick={onDeleteProject} color="inherit">
          <DeleteIcon />
        </IconButton>
      </div>

      <AddTask projectId={project.id} addItem={onChangeInputAdd} />

      <Typography variant="h6" color="inherit">
        To Do
      </Typography>
      <List>
        <ListItem key={'todo'} dense button>
          <ListItemText
            primary={' '}
            style={{
              marginRight: 150,
            }}
          />
          <ListItemSecondaryAction>
            <Typography
              variant="inherit"
              color="inherit"
              style={{
                marginRight: 25,
              }}
            >
              {'To finish at'}
            </Typography>
            <Typography
              variant="inherit"
              color="inherit"
              style={{
                marginRight: 50,
              }}
            >
              {'Created at'}
            </Typography>
          </ListItemSecondaryAction>
        </ListItem>
        {projectState && todoList.length > 0 ? (
          todoList.map((task: Task) => (
            <TaskItem
              task={task}
              changeStatus={handleStatusChange}
              deleteTask={handlerDelete}
              editTask={handlerEditTask}
            />
          ))
        ) : (
          <ListItem key="todo-empyt" dense button>
            <ListItemText primary="Empty list" />
          </ListItem>
        )}
      </List>

      <Typography variant="h6" color="inherit">
        Done
      </Typography>
      <List>
        {projectState && doneList.length > 0 ? (
          doneList.map((task: Task) => (
            <TaskItem
              task={task}
              changeStatus={handleStatusChange}
              deleteTask={handlerDelete}
              editTask={handlerEditTask}
            />
          ))
        ) : (
          <ListItem key="done-empyt" dense button>
            <ListItemText primary="Empty list" />
          </ListItem>
        )}
      </List>
    </main>
  );
};

export default ProjectDetails;
