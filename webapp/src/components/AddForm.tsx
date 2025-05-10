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

  return (
    <Form.Root onSubmit={handleSubmit}>
      <Form.Field name="title">
        <Form.Label htmlFor="title">title</Form.Label>
        <Form.Control type="text" required />
        <Form.Message match="valueMissing">required</Form.Message>
      </Form.Field>
      <Form.Submit asChild>
        <button
          type="submit"
          disabled={false} // TODO: hook this up
        >
          Add
        </button>
      </Form.Submit>
    </Form.Root>
  );
};
