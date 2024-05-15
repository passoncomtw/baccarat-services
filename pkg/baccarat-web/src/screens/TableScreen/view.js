import Grid from "@mui/material/Grid";
import Stack from "@mui/material/Stack";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import LoopIcon from "@mui/icons-material/Loop";
import EditIcon from "@mui/icons-material/Edit";
import { deepOrange } from "@mui/material/colors";

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
        <Grid container>
          <Grid item xs={12}>
            <Grid container>
                <Grid item xs={6} style={{backgroundColor: "red", height: 300}}>
                    <Stack direction="row" spacing={2}>
                        <div style={{height: 100, width: 80, backgroundColor: "#CCC"}}></div>
                        <div>莊</div>
                        <div style={{height: 100, width: 80, backgroundColor: "#CCC"}}></div>
                    </Stack>
                    <Stack direction="row" spacing={2}>
                        <div style={{height: 100, width: 80, backgroundColor: "#CCC"}}></div>
                        <div>閒</div>
                        <div style={{height: 100, width: 80, backgroundColor: "#CCC"}}></div>
                    </Stack>
                </Grid>
                <Grid item xs={3} style={{backgroundColor: "blue", height: 300}}></Grid>
                <Grid item xs={3} style={{backgroundColor: "grey", height: 300}}></Grid>
            </Grid>
          </Grid>
          <Grid item xs={12}>
            <Grid container spacing={2}>
              <Grid
                item
                xs={3}
                style={{
                  height: 250,
                  backgroundColor: "#CCCCCC",
                  margin: "20px",
                }}
              >
                <Typography>(1:0.95)</Typography>
                <Typography>莊</Typography>
              </Grid>
              <Grid
                item
                xs={3}
                style={{
                  height: 250,
                  backgroundColor: "#CCCCCC",
                  margin: "20px",
                }}
              >
                <Typography>(1:1)</Typography>
                <Typography>和</Typography>
              </Grid>
              <Grid
                item
                xs={3}
                style={{
                  height: 250,
                  backgroundColor: "#CCCCCC",
                  margin: "20px",
                }}
              >
                <Typography>(1:1)</Typography>
                <Typography>閒</Typography>
              </Grid>
            </Grid>
          </Grid>
          <Grid item xs={12}>
            <Stack direction="row" spacing={2}>
              <Button
                variant="contained"
                color="success"
                style={{
                  color: "#FFFFFF",
                  width: 60,
                  height: 60,
                  borderRadius: 60,
                }}
              >
                25
              </Button>
              <Button
                variant="contained"
                color="warning"
                style={{
                  color: "#FFFFFF",
                  width: 60,
                  height: 60,
                  borderRadius: 60,
                }}
              >
                50
              </Button>
              <Button
                variant="contained"
                color="error"
                style={{
                  color: "#FFFFFF",
                  width: 60,
                  height: 60,
                  borderRadius: 60,
                }}
              >
                100
              </Button>
              <Button
                variant="contained"
                color="primary"
                style={{
                  color: "#FFFFFF",
                  width: 60,
                  height: 60,
                  borderRadius: 60,
                }}
              >
                500
              </Button>
              <Button
                variant="contained"
                color="secondary"
                style={{
                  color: "#FFFFFF",
                  width: 60,
                  height: 60,
                  borderRadius: 60,
                }}
              >
                1000
              </Button>
              <Button
                variant="contained"
                color="info"
                style={{
                  color: "#FFFFFF",
                  width: 60,
                  height: 60,
                  borderRadius: 60,
                }}
              >
                5000
              </Button>
            </Stack>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
};

export default TableScreen;
