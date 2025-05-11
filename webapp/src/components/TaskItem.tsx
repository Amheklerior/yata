import type { FC } from "react";
import type { Task } from "../lib/types";
import { useDeleteTask, useUpdateTask } from "../lib/query";
import clsx from "clsx";

export const TaskItem: FC<{ task: Task }> = ({ task }) => {
  const { mutate: updateTask } = useUpdateTask(task.id);
  const { mutate: deleteTask } = useDeleteTask(task.id);

  const isComplete = task.status === "done";

  const handleUpdateTask = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();

    updateTask(
      { status: !isComplete ? "done" : "todo" },
      {
        onError: (error) => {
          // Give feedback to the user
          console.error(error);
        },
        onSuccess: () => {
          // TODO: Give feedback to the user
          console.log("Task updated");
        },
      },
    );
  };

  const handleDelete = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();

    deleteTask(undefined, {
      onError: (error) => {
        // Give feedback to the user
        console.error(error);
      },
      onSuccess: () => {
        // TODO: Give feedback to the user
        console.log("Task deleted");
      },
    });
  };

  return (
    <li className="flex items-center gap-4 p-4">
      <button onClick={handleUpdateTask} className="status-checkbox">
        <span className="sr-only">
          Mark as {isComplete ? "complete" : "to do"}
        </span>
        <div
          className={clsx(
            "inner-circle",
            isComplete ? "visible size-3" : "hidden",
          )}
        />
      </button>
      <h3 className={clsx("grow text-left", isComplete ? "line-through" : "")}>
        {task.title}
      </h3>
      <button
        onClick={handleDelete}
        className="font-mono text-lg hover:cursor-pointer hover:text-red-400"
      >
        x
      </button>
    </li>
  );
};
