
ALTER TABLE todoapp.users 
ADD COLUMN email 
VARCHAR(100) 
UNIQUE 
CHECK (
    email ~* '^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$' 
    AND email = LOWER(email)
);
