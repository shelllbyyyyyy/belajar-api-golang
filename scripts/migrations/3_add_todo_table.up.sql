DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'todo_status') THEN
        CREATE TYPE public.todo_status AS ENUM ('pending', 'working', 'paused', 'completed');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS public.todos (
    id uuid PRIMARY KEY NOT NULL,
    user_id uuid NOT NULL,
    name VARCHAR(100) NOT NULL,
    status public.todo_status DEFAULT 'pending',
    is_archived BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    started_at TIMESTAMP WITH TIME ZONE,
    paused_at TIMESTAMP WITH TIME ZONE,
    finished_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES public.users(id)
);
    
