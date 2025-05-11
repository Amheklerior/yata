import { use, useCallback } from "react";
import {
  Root as FormRoot,
  Field as FormField,
  Label as FormLabel,
  Control as FormControl,
  Message as FormMessage,
  Submit as FormSubmit,
} from "@radix-ui/react-form";
import clsx from "clsx/lite";
import addSound from "../assets/add-sound.wav";
import { NotificationCtx } from "../contexts/notificationCtx";
import { useCreateTask } from "../lib/query";
import { play } from "../lib/sounds";
import { Spinner } from "./Spinner";

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
    <FormRoot onSubmit={handleSubmit} className="flex items-start gap-4 p-4">
      <FormField name="title" className="flex grow flex-col gap-2">
        <FormLabel className="sr-only">title</FormLabel>
        <div className="flex items-center">
          <FormControl
            type="text"
            required
            placeholder="Add a new task..."
            className="input-field no-ring interactive animated grow"
          />
          <span
            className={clsx(
              "invisible relative right-9 flex h-1 w-1 items-center",
              isPending && "visible",
            )}
          >
            <Spinner />
          </span>
        </div>
        <div className="text-danger-400 min-h-6 text-left">
          <FormMessage match="valueMissing">
            <small>Are you joking?</small>
          </FormMessage>
        </div>
      </FormField>
      <FormSubmit asChild>
        <button type="submit" className="sr-only">
          Add
        </button>
      </FormSubmit>
    </FormRoot>
  );
};
