import { useState, useEffect } from "react";
import useCountDown from 'react-countdown-hook';
import Grid from "@mui/material/Grid";
import Stack from "@mui/material/Stack";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import LoopIcon from "@mui/icons-material/Loop";
import EditIcon from "@mui/icons-material/Edit";
import { deepOrange } from "@mui/material/colors";
import Chip from "./components/Chip";
import BettingArea from "./components/BettingArea";
import ResultArea from "./components/ResultArea";

const initialTime = 10 * 1000; // initial time in milliseconds, defaults to 60000
const interval = 1000; // interval to change remaining time amount, defaults to 1000

const styles = {
  avatarStyle: { bgcolor: deepOrange[500] },
};

const TableScreen = () => {
  const [step, setStep] = useState("Ready");
  const [timeLeft, { start, pause, resume, reset }] = useCountDown(initialTime, interval);

  useEffect(() => {
    start();
    return reset;
  }, []);

  useEffect(() => {
    if (timeLeft === 0) {
      if (step === "Ready") {
        setStep("Betting");
        start(initialTime);
      }

      if (step === "Betting") {
        setStep("StopBetting");
      }
      
    }
  },  [timeLeft, step, setStep]);
  return (
    <Grid container>
      <Grid item xs={2}>
        <Stack direction="column" spacing={2}>
          <Avatar sx={styles.avatarStyle}>N</Avatar>
          <Grid container>
            <Grid item>
              <Typography>tomas001</Typography>
            </Grid>
            <IconButton size="small">
              <EditIcon />
            </IconButton>
          </Grid>
          <Grid container>
            <Grid item>
              <Typography>額度: 1000</Typography>
            </Grid>
            <IconButton size="small">
              <LoopIcon />
            </IconButton>
          </Grid>
          <Grid container>
            <Grid item>
              <Button variant="contained">登出</Button>
            </Grid>
          </Grid>
        </Stack>
      </Grid>
      <Grid xs={10} item>
        <Grid container spacing={1}>
          <Grid item xs={12}>
            <ResultArea timeLeft={timeLeft} step={step} />
          </Grid>
          <Grid item xs={12}>
            <BettingArea items={[{odd: "1:0.95", text: "莊"}, {odd: "1:1", text: "和"}, {odd: "1:1", text: "閒"}]}/>
          </Grid>
          <Grid item xs={12}>
            <Stack direction="row" spacing={2} justifyContent="center">
              <Chip value={25} />
              <Chip value={50} />
              <Chip value={100} />
              <Chip value={500} />
              <Chip value={1000} />
              <Chip value={5000} />
            </Stack>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
};

export default TableScreen;
