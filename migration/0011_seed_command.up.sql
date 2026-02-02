INSERT INTO commands (name, description, usage, help_text) VALUES
(
  'help',
  'Displays available commands',
  'help [command]',
  'For more information on a specific command, type HELP command-name.'
),
(
  'whoami',
  'Displays information about the current user.',
  'whoami',
  'Shows name, role, location and description.'
),
(
  'skill',
  'Lists technical skills',
  'skills',
  'Lists technical/soft skills.'
),
(
  'education',
  'Lists education history',
  'education',
  'Displays education background.'
),
(
  'work',
  'Lists work experience',
  'work [name]',
  'Displays work experience history.'
),
(
  'showcase',
  'Displays portfolio projects',
  'showcase [id]',
  'Displays featured projects.'
);
