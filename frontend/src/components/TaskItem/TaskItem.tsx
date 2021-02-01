import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import { Task, UpdateTask } from '../../models/task';
import Checkbox from '@material-ui/core/Checkbox';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import IconButton from '@material-ui/core/IconButton';
import DeleteIcon from '@material-ui/icons/Delete';
import Typography from '@material-ui/core/Typography';
import EditTask from '../EditTask';

type TaskProps = {
  task: Task;
  changeStatus: (id: string, status: boolean) => void;
  editTask: (task: UpdateTask) => void;
  deleteTask: (id: string) => void;
};

const TaskItem: React.FC<TaskProps> = ({ task, changeStatus, deleteTask, editTask }: TaskProps) => {
  const [status, setTaskStatus] = React.useState(task.status);

  const [isEdit, setIsEdit] = React.useState(false);

  React.useEffect(() => {
    setTaskStatus(task.status);
  }, [task]);

  const handleStatusChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    if (task.status) {
      return;
    }
    const checked = e.target.checked;

    task.status = checked;
    setTaskStatus(checked);

    changeStatus(task.id, checked);
  };

  const handleDeleteClick = () => {
    if (task.status) {
      return;
    }
    deleteTask(task.id);
  };

  const timeToDate = (time: number): string => {
    return new Date(time * 1000).toLocaleDateString();
  };

  const onClickEditTask = () => {
    if (!task.status) {
      setIsEdit(true);
    }
  };

  const handlerUpdateTask = (task: Task | undefined) => {
    setIsEdit(false);

    if (!task) {
      return;
    }

    editTask({
      id: task.id,
      description: task.description,
      toFinishAt: task.toFinishAt,
    });
  };

  return (
    <ListItem key={task.id} dense button>
      {isEdit && <EditTask task={task} handlerUpdateTask={handlerUpdateTask} />}
      <Checkbox tabIndex={-1} disableRipple onChange={handleStatusChange} checked={status} />
      <ListItemText
        onClick={onClickEditTask}
        primary={task.description}
        style={{
          marginRight: 200,
        }}
      />
      <ListItemSecondaryAction>
        <Typography
          variant="inherit"
          color="inherit"
          style={{
            marginRight: 20,
          }}
        >
          {!task.status
            ? task.toFinishAt > 0
              ? timeToDate(task.toFinishAt)
              : '          '
            : task.finishedAt > 0
            ? timeToDate(task.finishedAt)
            : '          '}
        </Typography>
        <Typography variant="inherit" color="inherit">
          {timeToDate(task.createdAt)}
        </Typography>
        {!task.status && (
          <IconButton aria-label="Delete" onClick={handleDeleteClick}>
            <DeleteIcon />
          </IconButton>
        )}
      </ListItemSecondaryAction>
    </ListItem>
  );
};

export default TaskItem;
