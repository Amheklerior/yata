import { Form } from "radix-ui";
import { useCreateTask } from "../lib/query";
import { use, useCallback } from "react";
import clsx from "clsx";
import addSound from "../assets/add-sound.wav";
import { play } from "../lib/sounds";
import { Spinner } from "./Spinner";
import { NotificationCtx } from "../contexts/notificationCtx";

export const AddForm = () => {
  const { mutate: createTask, isPending } = useCreateTask();

  const { notify } = use(NotificationCtx);

  const handleSubmit = useCallback(
    (e: React.FormEvent<HTMLFormElement>) => {
      e.preventDefault();
      const formData = new FormData(e.currentTarget);
      const title = formData.get("title") as string;

      createTask(
        { title },
        {
          onError: () => notify("There was an error creating the task"),
          onSuccess: () => play(addSound),
        },
      );

      e.currentTarget.reset();
    },
    [createTask, notify],
  );

  return (
    <Form.Root onSubmit={handleSubmit} className="flex items-start gap-4 p-4">
      <Form.Field name="title" className="flex grow flex-col gap-2">
        <Form.Label className="sr-only">title</Form.Label>
        <div className="flex items-center">
          <Form.Control
            type="text"
            required
            placeholder="Add a new task..."
            className={clsx(
              "grow rounded-lg border px-3 py-2 placeholder:text-sm placeholder:uppercase",
              "transition-all duration-300",
              "hover:bg-stone-600/10 hover:placeholder:text-amber-200/30 focus:bg-stone-600/10 focus:placeholder:text-amber-200/30",
              "border-stone-400 caret-amber-200 placeholder:text-stone-400/70",
              "hover:border-amber-100 hover:ring-0 focus:border-amber-100 focus:ring-0 focus:outline-none",
            )}
          />
          <span
            className={clsx(
              "invisible relative right-9 flex h-1 w-0 items-center",
              isPending && "visible",
            )}
          >
            <Spinner />
          </span>
        </div>
        <div className="min-h-6 text-left text-red-400">
          <Form.Message match="valueMissing">
            <small>Are you joking?</small>
          </Form.Message>
        </div>
      </Form.Field>
      <Form.Submit asChild>
        <button type="submit" className="sr-only">
          Add
        </button>
      </Form.Submit>
    </Form.Root>
  );
};
