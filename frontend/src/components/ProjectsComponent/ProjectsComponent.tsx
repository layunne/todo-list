import React from 'react';

import { createStyles, Theme, makeStyles } from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import Divider from '@material-ui/core/Divider';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import AddIcon from '@material-ui/icons/Add';

import { Project } from '../../models/project';
import AddEditProject from '../AddEditProject';

const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    drawer: {
      width: drawerWidth,
      flexShrink: 0,
    },
    drawerPaper: {
      width: drawerWidth,
    },
    toolbar: theme.mixins.toolbar,
  }),
);

type ProjectsProps = {
  projects: Project[];
  updateSelectedProject: (project: Project) => void;
};

const ProjectsComponent: React.FC<ProjectsProps> = ({ projects, updateSelectedProject }: ProjectsProps) => {
  const classes = useStyles();

  const [isCreateProject, setIsCreateProject] = React.useState(false);

  const handleProjectClick = (project: Project) => {
    updateSelectedProject(project);
  };

  const handleCreateProject = (project: Project | undefined) => {
    setIsCreateProject(false);
    if (project) {
      handleProjectClick(project);
    }
  };

  return (
    <Drawer
      className={classes.drawer}
      variant="permanent"
      classes={{
        paper: classes.drawerPaper,
      }}
      anchor="left"
    >
      <div className={classes.toolbar} />

      {isCreateProject && <AddEditProject project={undefined} handlerUpdateProject={handleCreateProject} />}
      <List>
        <ListItem button onClick={() => setIsCreateProject(true)}>
          <ListItemIcon>
            <AddIcon />{' '}
          </ListItemIcon>
          <ListItemText primary="New Project" />
        </ListItem>
      </List>
      <Divider />
      <Divider />
      <Divider />
      <Divider />
      <Divider />
      <List>
        {projects.map((project) => (
          <ListItem button key={project.id} onClick={() => handleProjectClick(project)}>
            <ListItemText primary={project.name} />
          </ListItem>
        ))}
      </List>
    </Drawer>
  );
};

export default ProjectsComponent;
