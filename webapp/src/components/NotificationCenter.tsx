import { use, useEffect, useState } from "react";
import {
  Provider as ToastProvider,
  Root as ToastRoot,
  Description as ToastDescription,
  Viewport as ToastViewport,
} from "@radix-ui/react-toast";
import { NotificationCtx } from "../contexts/notificationCtx";

export const NotificationCenter = () => {
  const { message, clear } = use(NotificationCtx);
  const [open, setOpen] = useState(false);

  useEffect(() => {
    if (message) setOpen(true);
  }, [message]);

  const clearNotification = () => {
    setOpen(false);
    clear();
  };

  // TODO: fix enter and exit anim not working ...

  return (
    <ToastProvider duration={3000}>
      <ToastRoot
        open={open}
        onOpenChange={clearNotification}
        className="bg-danger-400 flex items-center rounded-lg px-4 py-3 shadow-lg"
      >
        <ToastDescription className="text text-danger-950">
          {message}
        </ToastDescription>
      </ToastRoot>

      <ToastViewport className="fixed bottom-0 right-0 z-[2147483647] m-0 flex w-[390px] max-w-[100vw] list-none flex-col gap-2.5 p-[var(--viewport-padding)] outline-none [--viewport-padding:_25px]" />
    </ToastProvider>
  );
};
