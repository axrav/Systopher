import React from "react";
import { Notification } from "@mantine/core";
import { IconAlertCircle } from "@tabler/icons";

function Error({
  error,
  showError,
  setShowError,
}: {
  error: string;
  showError: boolean;
  setShowError: React.Dispatch<React.SetStateAction<boolean>>;
}) {
  return (
    <div className="absolute bottom-5 right-5">
      <Notification
        icon={
          <IconAlertCircle
            className="bg-red-700 text-center rounded-full h-full w-full"
            size={18}
          />
        }
        onClose={() => setShowError(false)}
        title={<h3 className="text-lg font-bold">Login Error!</h3>}
        radius="md"
        color="red"
        className="bg-red-700"
        classNames={{
          closeButton: "hover:bg-red-500",
        }}
        hidden={!showError}
      >
        <p className="text-base max-w-xl text-white">{error}</p>
      </Notification>
    </div>
  );
}

export default Error;
