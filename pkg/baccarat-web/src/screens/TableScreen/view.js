import Grid from "@mui/material/Grid";
import Stack from "@mui/material/Stack";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import LoopIcon from "@mui/icons-material/Loop";
import EditIcon from "@mui/icons-material/Edit";
import { deepOrange } from "@mui/material/colors";
import HiddenCard from "screens/PokerScreen/components/cards/HiddenCard";
import Chip from "./components/Chip";
import BettingArea from "./components/BettingArea";

const TableScreen = () => {
  return (
    <Grid container>
      <Grid item xs={2}>
        <Stack direction="column" spacing={2}>
          <Avatar sx={{ bgcolor: deepOrange[500] }}>N</Avatar>
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
            <Grid container>
              <Grid item xs={4} style={{ height: 450 }}>
                <Stack direction="row" spacing={2} style={{ margin: 10 }}>
                  <div
                    style={{ height: 200, width: 160, backgroundColor: "#CCC" }}
                  ></div>
                  <div
                    style={{
                      display: "flex",
                      justifyContent: "center",
                      alignItems: "center",
                    }}
                  >
                    <div>莊</div>
                  </div>
                  <div
                    style={{ height: 200, width: 160, backgroundColor: "#CCC" }}
                  ></div>
                </Stack>
                <Stack direction="row" spacing={2} style={{ margin: 10 }}>
                  <div
                    style={{ height: 200, width: 160, backgroundColor: "#CCC" }}
                  ></div>
                  <div>閒</div>
                  <div
                    style={{ height: 200, width: 160, backgroundColor: "#CCC" }}
                  ></div>
                </Stack>
              </Grid>
              <div style={{ display: "flex", flex: 1, alignItems: "center" }}>
                <HiddenCard containerStyle={{ width: 100, height: 160 }} />
              </div>
              <Grid item xs={3} flex={2}>
                <Typography>準備投注</Typography>
                <Typography>10</Typography>
              </Grid>
            </Grid>
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
