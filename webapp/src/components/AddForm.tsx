import { Form } from "radix-ui";
import { useCreateTask } from "../lib/query";
import { useCallback } from "react";

export const AddForm = () => {
  const { mutate: createTask } = useCreateTask();

  const handleSubmit = useCallback(
    (e: React.FormEvent<HTMLFormElement>) => {
      e.preventDefault();
      const formData = new FormData(e.currentTarget);
      const title = formData.get("title") as string;

      // TODO: Add validation

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

  // TODO: Add validation
  // TODO: Add loading UI (while performing the submit)
  // TODO: prevent CLS when validation error appears

  return (
    <Form.Root
      onSubmit={handleSubmit}
      className="debug flex items-start gap-4 p-4"
    >
      <Form.Field name="title" className="flex grow flex-col gap-2">
        <Form.Label htmlFor="title" className="sr-only">
          title
        </Form.Label>
        <Form.Control
          type="text"
          required
          className="debug grow p-2 placeholder:font-light placeholder:text-gray-600"
          placeholder="Add a new task..."
        />
        <Form.Message
          match="valueMissing"
          className="text-left valid:invisible invalid:visible"
        >
          required
        </Form.Message>
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
