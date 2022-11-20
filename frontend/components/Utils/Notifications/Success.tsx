import React from "react";
import { IconCheck } from "@tabler/icons";
import { Notification } from "@mantine/core";

function Success({
  message,
  showSuccess,
  setShowSuccess,
  heading,
  loading,
}: {
  message: any;
  showSuccess: boolean;
  setShowSuccess: React.Dispatch<React.SetStateAction<boolean>>;
  heading: string;
  loading: boolean;
}) {
  return (
    <div className="absolute bottom-5 right-5">
      <Notification
        icon={<IconCheck size={20} />}
        color="teal"
        className="bg-teal-900"
        radius="md"
        hidden={!showSuccess}
        loading={loading}
        onClose={() => setShowSuccess(false)}
        title={<h3 className="text-xl">{heading}</h3>}
      >
        <p className="text-lg">{message}</p>
      </Notification>
    </div>
  );
}

export default Success;
