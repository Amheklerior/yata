import type { FC } from "react";
import type { Task } from "../lib/types";
import { useDeleteTask, useUpdateTask } from "../lib/query";
import clsx from "clsx";
import closeSvg from "../assets/close.svg";

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
          className={clsx("inner-circle", isComplete ? "scale-100" : "scale-0")}
        />
      </button>
      <h3
        className={clsx(
          "grow text-left transition-colors duration-300",
          isComplete ? "text-stone-200/50 line-through" : "",
        )}
      >
        {task.title}
      </h3>
      <button onClick={handleDelete} className="hover:cursor-pointer">
        <img
          src={closeSvg}
          alt="Delete task"
          width={24}
          height={24}
          className="close-icon"
        />
      </button>
    </li>
  );
};
