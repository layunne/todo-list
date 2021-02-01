import React from 'react';

import Input from '@material-ui/core/Input';
import Button from '@material-ui/core/Button';
import TasksService from '../../services/tasksService';

type AddTaskProps = {
  addItem: () => void;
  projectId: string;
};

const AddTask: React.FC<AddTaskProps> = ({ addItem, projectId }: AddTaskProps) => {
  const [description, setDescription] = React.useState('');

  const onChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    setDescription(e.target.value);
  };

  const handlerAdd = async () => {
    console.log('handlerAdd');

    await TasksService.create({
      projectId,
      description,
      toFinishAt: 0,
    });
    setDescription('');
    addItem();
  };

  return (
    <div style={{ display: 'flex', marginBottom: '30px', height: '50px' }}>
      <Input
        placeholder="Description"
        inputProps={{
          'aria-label': 'Description',
        }}
        value={description}
        onChange={onChange}
        style={{ width: '90%' }}
      />

      <Button onClick={handlerAdd} variant="contained" color="primary" style={{ width: '10%' }}>
        Add
      </Button>
    </div>
  );
};

export default AddTask;
