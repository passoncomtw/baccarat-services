import { Navigate } from "react-router-dom";
import { Divider, Grid, Typography } from "@mui/material";

// project imports
import AuthWrapper1 from "components/AuthWrapper1";
import AuthCardWrapper from "components/AuthCardWrapper";
import AuthLogin from "components/auth-forms/AuthLogin";
import AuthFooter from "ui-component/cards/AuthFooter";
import packageJson from "../../../package.json";

const LoginScreen = (props) => {
  if (props.isAuth) return <Navigate to="/" replace />;
  return (
    <AuthWrapper1>
      <Grid
        container
        direction="column"
        justifyContent="flex-end"
        sx={{ minHeight: "100vh" }}
      >
        <Grid item xs={12}>
          <Grid
            container
            justifyContent="center"
            alignItems="center"
            sx={{ minHeight: "calc(100vh - 68px)" }}
          >
            <Grid item sx={{ m: { xs: 1, sm: 3 }, mb: 0 }}>
              <AuthCardWrapper>
                <Grid
                  container
                  spacing={2}
                  alignItems="center"
                  justifyContent="center"
                >                                    
                  <Grid item xs={12}>
                    <AuthLogin />
                  </Grid>
                  <Grid item xs={12}>
                    <Divider />
                  </Grid>
                  <Grid item xs={12}>
                    <Grid
                      item
                      container
                      direction="column"
                      alignItems="center"
                      xs={12}
                    >
                      <Typography
                        variant="subtitle1"
                        sx={{ textDecoration: "none" }}
                      >
                        V{packageJson.version}
                      </Typography>
                    </Grid>
                  </Grid>
                </Grid>
              </AuthCardWrapper>
            </Grid>
          </Grid>
        </Grid>
        <Grid item xs={12} sx={{ m: 3, mt: 1 }}>
          <AuthFooter />
        </Grid>
      </Grid>
    </AuthWrapper1>
  );
};

export default LoginScreen;
