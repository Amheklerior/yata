import { Form } from "radix-ui";
import { useCreateTask } from "../lib/query";
import { useCallback } from "react";
import clsx from "clsx";

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
            // Give feedback to the user
            console.error(error);
          },
          onSuccess: () => {
            // TODO: Give feedback to the user
          },
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
          className={clsx(
            "grow p-2",
            "rounded-lg border border-stone-400 caret-amber-200 ring-amber-200",
            "placeholder:text-stone-400/70",
            "hover:ring-1 focus:ring-1 focus:ring-amber-200 focus:outline-none",
          )}
          placeholder="Add a new task..."
        />
        <div className="min-h-6 text-left text-red-400">
          <Form.Message match="valueMissing">
            <small>Are you joking?</small>
          </Form.Message>
        </div>
      </Form.Field>
      <Form.Submit asChild>
        <button
          type="submit"
          className="sr-only"
          disabled={false} // TODO: hook this up
        >
          Add
        </button>
      </Form.Submit>
    </Form.Root>
  );
};
