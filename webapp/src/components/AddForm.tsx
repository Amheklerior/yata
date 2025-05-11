import { Form } from "radix-ui";
import { useCreateTask } from "../lib/query";
import { useCallback } from "react";
import clsx from "clsx";
import addSound from "../assets/add-sound.wav";
import { play } from "../lib/sounds";

export const AddForm = () => {
  const { mutate: createTask } = useCreateTask();

  const handleSubmit = useCallback(
    (e: React.FormEvent<HTMLFormElement>) => {
      e.preventDefault();
      const formData = new FormData(e.currentTarget);
      const title = formData.get("title") as string;

      createTask(
        { title },
        {
          onError: (error) => {
            // TODO: Give feedback to the user
            console.error(error);
          },
          onSuccess: () => play(addSound),
        },
      );

      e.currentTarget.reset();
    },
    [createTask],
  );

  // TODO: Add loading UI (while performing the submit)

  return (
    <Form.Root onSubmit={handleSubmit} className="flex items-start gap-4 p-4">
      <Form.Field name="title" className="flex grow flex-col gap-2">
        <Form.Label htmlFor="title" className="sr-only">
          title
        </Form.Label>
        <Form.Control
          type="text"
          required
          placeholder="Add a new task..."
          className={clsx(
            "grow rounded-lg border px-3 py-2 placeholder:text-sm placeholder:uppercase",
            "transition-all duration-300",
            "hover:bg-stone-600/10 focus:bg-stone-600/10",
            "border-stone-400 caret-amber-200 placeholder:text-stone-400/70",
            "hover:border-amber-100 hover:ring-0 focus:border-amber-100 focus:ring-0 focus:outline-none",
          )}
        />
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
