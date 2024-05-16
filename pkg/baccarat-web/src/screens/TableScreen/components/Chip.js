import Button from "@mui/material/Button";

const settings = {
  25: {
    color: "success",
    style: {
      color: "#FFFFFF",
      width: 60,
      height: 60,
      borderRadius: 60,
    },
  },
  50: {
    color: "warning",
    style: {
      color: "#FFFFFF",
      width: 60,
      height: 60,
      borderRadius: 60,
    },
  },
  100: {
    color: "error",
    style: {
      color: "#FFFFFF",
      width: 60,
      height: 60,
      borderRadius: 60,
    },
  },
  500: {
    color: "primary",
    style: {
      color: "#FFFFFF",
      width: 60,
      height: 60,
      borderRadius: 60,
    },
  },
  1000: {
    color: "secondary",
    style: {
      color: "#FFFFFF",
      width: 60,
      height: 60,
      borderRadius: 60,
    },
  },
  5000: {
    color: "info",
    style: {
      color: "#FFFFFF",
      width: 60,
      height: 60,
      borderRadius: 60,
    },
  },
};

const Chip = (props) => {
  const { value } = props;
  const setting = settings[value];

  return (
    <Button variant="contained" color={setting.color} style={setting.style}>
      {value}
    </Button>
  );
};

export default Chip;
