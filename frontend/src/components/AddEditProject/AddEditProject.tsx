import React from 'react';

import { Project } from '../../models/project';
import ProjectsService from '../../services/projectsService';
import Dialog from '@material-ui/core/Dialog';
import DialogTitle from '@material-ui/core/DialogTitle';
import DialogActions from '@material-ui/core/DialogActions';
import Button from '@material-ui/core/Button';
import DialogContent from '@material-ui/core/DialogContent';
import TextField from '@material-ui/core/TextField';

type ProjectProps = {
  project: Project | undefined;
  handlerUpdateProject: (project: Project | undefined) => void;
};

const AddEditProject: React.FC<ProjectProps> = ({ project, handlerUpdateProject }: ProjectProps) => {
  const [open, setOpen] = React.useState(true);
  const [description, setDescription] = React.useState(project ? project.name : '');

  const onChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    setDescription(e.target.value);
  };

  const handleSubmit = async () => {
    let p: Project;

    if (description.length === 0) {
      return;
    }

    setOpen(false);

    if (project) {
      const response = await ProjectsService.update({
        id: project.id,
        name: description,
      });
      p = response.data;
    } else {
      const response = await ProjectsService.create({
        name: description,
      });
      p = response.data;
    }
    handlerUpdateProject(p);
  };

  const handleClose = () => {
    setOpen(false);
    handlerUpdateProject(undefined);
  };

  return (
    <>
      <Dialog open={open} onClose={handleClose} aria-labelledby="form-dialog-title">
        <DialogTitle id="form-dialog-title">Project</DialogTitle>
        <DialogContent>
          <TextField
            autoFocus
            margin="dense"
            id="name"
            label="Project Name"
            variant="outlined"
            fullWidth
            onChange={onChange}
            value={description}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} color="primary">
            Cancel
          </Button>
          <Button onClick={handleSubmit} color="primary">
            {project ? 'Edit' : 'Create'}
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default AddEditProject;
