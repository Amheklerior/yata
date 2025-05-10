import { z } from "zod";

export const TaskSchema = z.object({
  id: z.number().min(1),
  title: z.string().min(1, "Title must be a non-empty string"),
  detail: z.string().optional(),
  status: z.union([z.literal("todo"), z.literal("done")]),
});

export const CreateTaskReqBodySchema = z.object({
  title: z.string().min(1, "Title must be a non-empty string"),
  detail: z.string().optional(),
});

export const UpdateTaskReqBodySchema = z.object({
  title: z.string().optional(),
  detail: z.string().optional(),
  status: z.union([z.literal("todo"), z.literal("done")]).optional(),
});

export const GetTasksResponseBodySchema = z.object({
  tasks: z.array(TaskSchema),
  total: z.number(),
});

export const ErrorResponseSchema = z.object({
  error: z.string(),
});

export const CreateTaskResponseBodySchema = z.object({
  task: TaskSchema,
});

export const GetTaskByIdResponseBodySchema = z.object({
  task: TaskSchema,
});

export const UpdateTaskResponseBodySchema = z.object({
  task: TaskSchema,
});

export const DeleteTaskResponseBodySchema = z.object({
  message: z.string(),
});
