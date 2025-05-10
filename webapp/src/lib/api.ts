import { HttpClient } from "./httpClient";
import { CreateTaskReqBodySchema, UpdateTaskReqBodySchema } from "./schemas";
import type {
  CreateTaskReqBody,
  CreateTaskResponse,
  DeleteTaskResponse,
  GetTaskByIdResponse,
  GetTasksResponse,
  UpdateTaskReqBody,
  UpdateTaskResponse,
} from "./types";

// TODO: handle all response and error cases

export const getTasks = async (): Promise<GetTasksResponse> => {
  return await HttpClient.get("tasks");
};

export const createTask = async (
  reqBody: CreateTaskReqBody,
): Promise<CreateTaskResponse> => {
  CreateTaskReqBodySchema.parse(reqBody);
  return await HttpClient.post("tasks", reqBody);
};

export const getTaskById = async (id: number): Promise<GetTaskByIdResponse> => {
  return await HttpClient.get(`tasks/${id}`);
};

export const updateTask = async (
  id: number,
  reqBody: UpdateTaskReqBody,
): Promise<UpdateTaskResponse> => {
  UpdateTaskReqBodySchema.parse(reqBody);
  return await HttpClient.put(`tasks/${id}`, reqBody);
};

export const deleteTask = async (id: number): Promise<DeleteTaskResponse> => {
  return await HttpClient.delete(`tasks/${id}`);
};
