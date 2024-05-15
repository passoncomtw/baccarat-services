import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Grid from "@mui/material/Grid";
import Stack from "@mui/material/Stack";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import LoopIcon from "@mui/icons-material/Loop";
import EditIcon from "@mui/icons-material/Edit";
import { deepOrange } from "@mui/material/colors";

const RoomScreen = () => {
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
        <Grid container spacing={2}>
          <Grid item>
            <Card sx={{ minWidth: 250 }}>
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  百家樂 A01
                </Typography>
              </CardContent>
              <CardActions>
                <Button size="small">進入遊戲</Button>
              </CardActions>
            </Card>
          </Grid>
          <Grid item>
            <Card sx={{ minWidth: 250 }}>
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  百家樂 A02
                </Typography>
              </CardContent>
              <CardActions>
                <Button size="small">進入遊戲</Button>
              </CardActions>
            </Card>
          </Grid>
          <Grid item>
            <Card sx={{ minWidth: 250 }}>
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  百家樂 A03
                </Typography>
              </CardContent>
              <CardActions>
                <Button size="small">進入遊戲</Button>
              </CardActions>
            </Card>
          </Grid>
          <Grid item>
            <Card sx={{ minWidth: 250 }}>
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  百家樂 A04
                </Typography>
              </CardContent>
              <CardActions>
                <Button size="small">進入遊戲</Button>
              </CardActions>
            </Card>
          </Grid>
          <Grid item>
            <Card sx={{ minWidth: 250 }}>
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  百家樂 A05
                </Typography>
              </CardContent>
              <CardActions>
                <Button size="small">進入遊戲</Button>
              </CardActions>
            </Card>
          </Grid>
          <Grid item>
            <Card sx={{ minWidth: 250 }}>
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  百家樂 A06
                </Typography>
              </CardContent>
              <CardActions>
                <Button size="small">進入遊戲</Button>
              </CardActions>
            </Card>
          </Grid>
          <Grid item>
            <Card sx={{ minWidth: 250 }}>
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  百家樂 A07
                </Typography>
              </CardContent>
              <CardActions>
                <Button size="small">進入遊戲</Button>
              </CardActions>
            </Card>
          </Grid>
        </Grid>
      </Grid>
    </Grid>
  );
};

export default RoomScreen;
