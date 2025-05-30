import { use, type FC } from "react";
import clsx from "clsx/lite";
import closeSvg from "../assets/close.svg";
import checkSound from "../assets/mark-sound.wav";
import deleteSound from "../assets/delete-sound.wav";
import { NotificationCtx } from "../contexts/notificationCtx";
import { useDeleteTask, useUpdateTask } from "../lib/query";
import { play } from "../lib/sounds";
import type { Task } from "../lib/types";

export const TaskItem: FC<{ task: Task }> = ({ task }) => {
  const { mutate: updateTask } = useUpdateTask(task.id);
  const { mutate: deleteTask } = useDeleteTask(task.id);

  const { notify } = use(NotificationCtx);

  const isComplete = task.status === "done";

  const handleUpdateTask = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();

    updateTask(
      { status: !isComplete ? "done" : "todo" },
      {
        onError: () => notify("There was an error updating the task"),
        onSuccess: () => play(checkSound),
      },
    );
  };

  const handleDelete = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();

    deleteTask(undefined, {
      onError: () => notify("There was an error deleting the task"),
      onSuccess: () => play(deleteSound),
    });
  };

  return (
    <li className="flex items-center gap-4 p-4">
      <button
        onClick={handleUpdateTask}
        className="status-checkbox interactive animated min-w-5"
        data-anim="slow"
      >
        <span className="sr-only">
          Mark as {isComplete ? "complete" : "to do"}
        </span>
        <div
          className={clsx("inner-circle", isComplete ? "scale-100" : "scale-0")}
        />
      </button>
      <h3
        className={clsx(
          "animated max-w-full grow overflow-x-hidden text-left text-nowrap text-ellipsis",
          isComplete ? "text-neutral-200/50 line-through" : "",
        )}
        data-anim="slow"
      >
        {task.title}
      </h3>
      <button
        onClick={handleDelete}
        className="interactive no-ring close-btn animated min-w-6"
        data-anim="slow"
      >
        <img src={closeSvg} alt="Delete task" width={24} height={24} />
      </button>
    </li>
  );
};
