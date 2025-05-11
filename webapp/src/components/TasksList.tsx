import type { FC } from "react";
import { TaskItem } from "./TaskItem";
import { useGetTasks } from "../lib/query";
import type { Task } from "../lib/types";
import { TaskItemSkeleton } from "./TaskItemScheleton";

const Loading: FC = () => (
  <div>
    {Array.from({ length: 4 }).map((_, i) => (
      <TaskItemSkeleton key={i} />
    ))}
  </div>
);

const Error: FC = () => <p>Error</p>;
const EmptyList: FC = () => (
  <p className="text-stone-300">No tasks Yet. Start adding some...</p>
);

export const TasksList: FC = () => {
  const { data, isLoading, isSuccess } = useGetTasks();

  if (isLoading) return <Loading />;

  if (!isSuccess) return <Error />;

  // TODO: blur top and bottom of the list if contines in that direction

  return (
    <>
      {data.total === 0 ? (
        <EmptyList />
      ) : (
        <ul className="no-scrollbar relative max-h-[420px] overflow-y-auto">
          {data.tasks.map((task: Task) => (
            <TaskItem key={task.id} task={task} />
          ))}
        </ul>
      )}
    </>
  );
};
