import { QueryClient, useMutation, useQuery } from "@tanstack/react-query";
import {
  createTask,
  deleteTask,
  getTaskById,
  getTasks,
  updateTask,
} from "./api";
import type { CreateTaskReqBody, UpdateTaskReqBody } from "./types";

export const queryClient = new QueryClient({});

export const useGetTasks = () => {
  return useQuery({
    queryKey: ["get-tasks"],
    queryFn: getTasks,
  });
};

export const useGetTaskById = (id: number) => {
  return useQuery({
    queryKey: ["get-task", id],
    queryFn: () => getTaskById(id),
    enabled: !!id,
  });
};

export const useCreateTask = () => {
  return useMutation({
    mutationKey: ["create-task"],
    mutationFn: (body: CreateTaskReqBody) => createTask(body),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["get-tasks"] });
    },
  });
};

export const useUpdateTask = (id: number) => {
  return useMutation({
    mutationKey: ["update-task", id],
    mutationFn: (body: UpdateTaskReqBody) => updateTask(id, body),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["get-tasks"] });
      queryClient.invalidateQueries({ queryKey: ["get-task", id] });
    },
  });
};

export const useDeleteTask = (id: number) => {
  return useMutation({
    mutationKey: ["delete-task", id],
    mutationFn: () => deleteTask(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["get-tasks"] });
      queryClient.invalidateQueries({ queryKey: ["get-task", id] });
    },
  });
};
