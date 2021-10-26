db = db.getSiblingDB('logger_fitness_db');
db.createUser(
{
  user: 'logger_fitness',
  pwd: 'lfPass', 
  roles: [
    { role: 'readWrite', db: 'logger_fitness_db' }
  ]
});
