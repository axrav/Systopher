import { Checkbox } from "@mantine/core";

function RememberMe({
  checked,
  setChecked,
}: {
  checked: boolean;
  setChecked: any;
}) {
  return (
    <div>
      <Checkbox
        color="blue"
        radius="md"
        checked={checked}
        onChange={(e) => setChecked(!checked)}
        size="md"
        description="To reduce the number or logins"
        label="Remember User on this device?"
      />
    </div>
  );
}

export default RememberMe;
