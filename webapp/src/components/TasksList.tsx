import type { FC } from "react";
import { TaskItem } from "./TaskItem";
import { useGetTasks } from "../lib/task/query";

const Loading: FC = () => <p>Loading...</p>;
const Error: FC = () => <p>Error</p>;
const EmptyList: FC = () => <p>No tasks Yet. Start adding some...</p>;

export const TasksList: FC = () => {
  const { data, isLoading, isSuccess } = useGetTasks();

  if (isLoading || !isSuccess) return <Loading />;

  if (!isSuccess) return <Error />;

  return (
    <section id="tasks-list">
      {data.total === 0 ? (
        <EmptyList />
      ) : (
        <ul>
          {data.tasks.map((task) => (
            <TaskItem key={task.id} task={task} />
          ))}
        </ul>
      )}
    </section>
  );
};
