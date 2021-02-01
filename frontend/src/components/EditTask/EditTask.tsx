import React from 'react';

import Dialog from '@material-ui/core/Dialog';
import DialogTitle from '@material-ui/core/DialogTitle';
import DialogActions from '@material-ui/core/DialogActions';
import Button from '@material-ui/core/Button';
import DialogContent from '@material-ui/core/DialogContent';
import TextField from '@material-ui/core/TextField';
import { Task } from '../../models/task';
import TasksService from '../../services/tasksService';

const dateToString = (date: Date) => {
  return date.toISOString().split('.')[0];
};

type EditTaskProps = {
  task: Task;
  handlerUpdateTask: (task: Task | undefined) => void;
};

const EditTask: React.FC<EditTaskProps> = ({ task, handlerUpdateTask }: EditTaskProps) => {
  const [open, setOpen] = React.useState(true);

  const [description, setDescription] = React.useState(task.description);
  const [toFinishAt, setToFinishAt] = React.useState(task.toFinishAt);

  const [selectedDate, setSelectedDate] = React.useState<string>(
    dateToString(task.toFinishAt > 0 ? new Date(task.toFinishAt * 1000) : new Date()),
  );

  const handleDateChange = (date: Date | null) => {
    if (date) {
      setToFinishAt(Math.round(date.getTime() / 1000));
    }
  };

  const onChangeDescription = async (e: React.ChangeEvent<HTMLInputElement>) => {
    setDescription(e.target.value);
  };

  const onChangeToFinishAt = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    const date = new Date(value);
    handleDateChange(date);
    setSelectedDate(value);
    setToFinishAt(Math.round(date.getTime() / 1000));
  };

  const handleSubmit = async () => {
    if (description.length === 0) {
      return;
    }

    const response = await TasksService.update({
      id: task.id,
      description,
      toFinishAt,
    });

    setOpen(false);

    handlerUpdateTask(response.data);
  };

  const handleClose = () => {
    setOpen(false);
    handlerUpdateTask(undefined);
  };

  return (
    <>
      <Dialog open={open} onClose={handleClose} aria-labelledby="form-dialog-title">
        <DialogTitle id="form-dialog-title">Task</DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            margin="dense"
            id="description"
            label="Description"
            fullWidth
            onChange={onChangeDescription}
            value={description}
          />
          <TextField
            id="datetime-local"
            label="To finish at"
            type="datetime-local"
            value={selectedDate}
            onChange={onChangeToFinishAt}
            InputLabelProps={{
              shrink: true,
            }}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} color="primary">
            Cancel
          </Button>
          <Button onClick={handleSubmit} color="primary">
            Edit
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default EditTask;
