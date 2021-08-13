import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import InputLabel from '@material-ui/core/InputLabel';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import DateFnsUtils from '@date-io/date-fns';
import { KeyboardDatePicker, MuiPickersUtilsProvider } from "@material-ui/pickers";

const useStyles = makeStyles((theme) => ({
  formControl: {
    margin: theme.spacing(0),
    minWidth: 265,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
}));

const AddressForm = ({ applicantFormValues, changeApplicantFormValues, dateOfBirth, setDateOfBirth, employment, setEmploymentType }) => {

  const handleDateChange = (date) => {
    setDateOfBirth(date);
  };

  const classes = useStyles();

  const handleChange = (event) => {
    const name = event.target.name;
    setEmploymentType({
      ...employment,
      [name]: event.target.value,
    });
  };

  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom>
        Applicant info
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} sm={6}>
          <TextField
            required
            id="firstName"
            name="firstName"
            label="First name"
            fullWidth
            value={applicantFormValues.firstName}
            onChange={(e) => changeApplicantFormValues('firstName', e.target.value)}
          />
        </Grid>
        <Grid item xs={12} sm={6}>
          <TextField
            required
            id="lastName"
            name="lastName"
            label="Last name"
            fullWidth
            value={applicantFormValues.lastName}
            onChange={(e) => changeApplicantFormValues('lastName', e.target.value)}
          />
        </Grid>
        <Grid item xs={12} sm={6}>
          <MuiPickersUtilsProvider utils={DateFnsUtils}>
            <KeyboardDatePicker
              required
              margin="normal"
              id="dateOfBirth"
              label="Date of birth"
              format="MM/dd/yyyy"
              value={dateOfBirth}
              onChange={handleDateChange}
              KeyboardButtonProps={{
                'aria-label': 'change date',
              }}
            />
          </MuiPickersUtilsProvider>
        </Grid>
        <Grid item xs={12}>
          <TextField
            required
            id="phoneNumber"
            name="phoneNumber"
            label="Phone Number"
            fullWidth
            value={applicantFormValues.phoneNumber}
            onChange={(e) => changeApplicantFormValues('phoneNumber', e.target.value)}
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            required
            id="socialSecurityNumber"
            name="socialSecurityNumber"
            label="Social Security Number"
            fullWidth
            value={applicantFormValues.socialSecurityNumber}
            onChange={(e) => changeApplicantFormValues('socialSecurityNumber', e.target.value)}
          />
        </Grid> 
        <Grid item xs={12} sm={6}>  
        <FormControl className={classes.formControl}>
            <InputLabel htmlFor="employmentType">Employment details</InputLabel>
            <Select
              native
              value={employment['employmentType']}
              onChange={handleChange}
              name="employmentType"
              inputProps={{
                id: 'employmentType',
              }}
            >
              <option aria-label="None" value="" />
              <option value="Employer">Employer</option>
              <option value="Previous Employer">Previous Employer</option>
            </Select>
          </FormControl>
          </Grid>
        <Grid item xs={12} sm={6}>
          <TextField
            required
            id="occupation"
            name="occupation"
            label="Occupation"
            fullWidth
            value={applicantFormValues.occupation}
            onChange={(e) => changeApplicantFormValues('occupation', e.target.value)}
          />
        </Grid>   
        <Grid item xs={12} sm={6}>
          <TextField
            required
            id="monthlyIncome"
            name="monthlyIncome"
            label="Monthly Income"
            fullWidth
            value={applicantFormValues.monthlyIncome}
            onChange={(e) => changeApplicantFormValues('monthlyIncome', e.target.value)}
          />
        </Grid>   
        <Grid item xs={12} sm={6}>
          <TextField
            required
            id="yearsEmployed"
            type="number"
            name="yearsEmployed"
            label="Years Employed"
            fullWidth
            value={applicantFormValues.yearsEmployed}
            onChange={(e) => changeApplicantFormValues('yearsEmployed', e.target.value)}
          />
        </Grid>  
      </Grid>
    </React.Fragment>
  );
}
export default AddressForm;