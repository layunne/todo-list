/* eslint-disable react-hooks/exhaustive-deps */
import React, { useState, useEffect } from 'react';
import Navbar from '../../components/Navbar';

import { createStyles, Theme, makeStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import ProjectsComponent from '../../components/ProjectsComponent';
import ProjectDetails from '../../components/ProjectDetails';
import ProjectsService from '../../services/projectsService';
import { Project } from '../../models/project';

const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
    },
    appBar: {
      width: `calc(100% - ${drawerWidth}px)`,
      marginLeft: drawerWidth,
    },
    drawer: {
      width: drawerWidth,
      flexShrink: 0,
    },
    drawerPaper: {
      width: drawerWidth,
    },
    // necessary for content to be below app bar
    toolbar: theme.mixins.toolbar,
    content: {
      flexGrow: 1,
      backgroundColor: theme.palette.background.default,
      padding: theme.spacing(3),
    },
  }),
);

const Home: React.FC = () => {
  const classes = useStyles();

  const [projects, setProjects] = useState<Project[]>([]);

  const [selectedProject, setSelectedProject] = useState<Project>();

  useEffect(() => {
    getProjects();
  }, []);

  const getProjects = async () => {
    const response = await ProjectsService.getAll();
    const projects: Project[] = response.data;
    setProjects(projects);

    if (!selectedProject && projects.length > 0) {
      setSelectedProject(projects[0]);
    } else if (selectedProject) {
      const index = projects.findIndex((p: Project) => p.id === selectedProject.id);

      if (index < 0) {
        if (projects.length > 0) {
          setSelectedProject(projects[0]);
        } else {
          setSelectedProject(undefined);
        }
      } else {
        setSelectedProject(projects[index]);
      }
    }
  };

  const updateSelectedProject = async (project: Project) => {
    const index = projects.findIndex((p: Project) => p.id === project.id);

    if (index >= 0) {
      setSelectedProject({
        id: project.id,
        name: project.name,
        userId: project.userId,
        tasks: project.tasks,
      });
    }

    const response = await ProjectsService.getAll();
    const ps: Project[] = response.data;
    setProjects(ps);
  };

  const handlerUpdateProject = async (project: Project | undefined) => {
    if (project) {
      updateSelectedProject(project);
      return;
    } else {
      getProjects();
    }
  };

  return (
    <div className={classes.root}>
      <CssBaseline />
      <Navbar />
      <ProjectsComponent projects={projects} updateSelectedProject={updateSelectedProject} />
      {selectedProject && <ProjectDetails project={selectedProject} handlerUpdateProject={handlerUpdateProject} />}
    </div>
  );
};

export default Home;
