import type { z } from "zod";
import type {
  CreateTaskReqBodySchema,
  CreateTaskResponseBodySchema,
  DeleteTaskResponseBodySchema,
  ErrorResponseSchema,
  GetTaskByIdResponseBodySchema,
  GetTasksResponseBodySchema,
  TaskSchema,
  UpdateTaskReqBodySchema,
  UpdateTaskResponseBodySchema,
} from "./schemas";

// Task object type
export type Task = z.infer<typeof TaskSchema>;

// Request types
export type CreateTaskReqBody = z.infer<typeof CreateTaskReqBodySchema>;
export type UpdateTaskReqBody = z.infer<typeof UpdateTaskReqBodySchema>;

// Responses types
export type ErrorResponse = z.infer<typeof ErrorResponseSchema>;
export type GetTasksResponse = z.infer<typeof GetTasksResponseBodySchema>;
export type CreateTaskResponse = z.infer<typeof CreateTaskResponseBodySchema>;
export type GetTaskByIdResponse = z.infer<typeof GetTaskByIdResponseBodySchema>;
export type UpdateTaskResponse = z.infer<typeof UpdateTaskResponseBodySchema>;
export type DeleteTaskResponse = z.infer<typeof DeleteTaskResponseBodySchema>;
