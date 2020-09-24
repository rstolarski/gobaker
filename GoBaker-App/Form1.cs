using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Diagnostics;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace GoBaker_App
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void Form1_Load(object sender, EventArgs e)
        {

        }

        private void button1_Click(object sender, EventArgs e)
        {
            Console.WriteLine("READING LOW");
            OpenFileDialog ofd = new OpenFileDialog();
            if (ofd.ShowDialog() == System.Windows.Forms.DialogResult.OK)
            {
                Program.LowPolyFile = ofd.FileName;
            }
                
        }

        private void button2_Click(object sender, EventArgs e)
        {
            Console.WriteLine("READING HIGH");
            OpenFileDialog ofd = new OpenFileDialog();
            if (ofd.ShowDialog() == System.Windows.Forms.DialogResult.OK)
            {
                Program.HighPolyFile = ofd.FileName;
            }
        }

        private void button3_Click(object sender, EventArgs e)
        {
            Console.WriteLine("READING PLY");
            OpenFileDialog ofd = new OpenFileDialog();
            if (ofd.ShowDialog() == System.Windows.Forms.DialogResult.OK)
            {
                Program.HighPolyPLYFile = ofd.FileName;
            }
        }

        private void button4_Click(object sender, EventArgs e)
        {
            Console.WriteLine("BAKING");


            if (Program.LowPolyFile == "")
            {
                MessageBox.Show("You did not set path to Lowpoly file. Double check again", "Baker error");
            }
            else if (Program.HighPolyFile == "")
            {
                MessageBox.Show("You did not set path to Highpoly file. Double check again", "Baker error");
            }
            else if (Program.Output == "")
            {
                MessageBox.Show("You did not set output path. Double check again", "Baker error");
            }
            else if (Program.RenderSize == "")
            {
                MessageBox.Show("You did not set render size. Double check again", "Baker error");
            }
            else if (Program.MaxFrontDistance == "")
            {
                MessageBox.Show("You did not set max frontal distance. Double check again", "Baker error");
            }
            else if (Program.MaxRearDistance == "")
            {
                MessageBox.Show("You did not set max rear distance Double check again", "Baker error");
            }
            else if (Program.ReadIDMap && Program.HighPolyPLYFile == "")
            {
                MessageBox.Show("You did not set path to PLY file. Double check again", "Baker error");
            }
            else
            {
                Process p = new Process();
                p.StartInfo.FileName = @"gobaker.exe";

                string arguments = "";
                arguments += " -l \"" + Program.LowPolyFile + "\"";
                arguments += " -h \"" + Program.HighPolyFile + "\"";
                arguments += " -hp \"" + Program.HighPolyPLYFile + "\"";
                arguments += " -s " + Program.RenderSize;
                arguments += " -frontD " + Program.MaxFrontDistance;
                arguments += " -rearD " + Program.MaxRearDistance;
                arguments += " -o \"" + Program.Output + "\"";
                arguments += " -id=" + Program.ReadIDMap;

                Console.WriteLine(arguments);

                p.StartInfo.Arguments = arguments;
                p.StartInfo.UseShellExecute = false;
                p.StartInfo.RedirectStandardOutput = true;
                p.StartInfo.WindowStyle = System.Diagnostics.ProcessWindowStyle.Hidden;
                p.StartInfo.CreateNoWindow = false; //not diplay a windows
                p.Start();
                string output = p.StandardOutput.ReadToEnd(); //The output result
                p.WaitForExit();
            }
        }

        private void label1_Click(object sender, EventArgs e)
        {
        }

        private void renderSizeBox_TextChanged(object sender, EventArgs e)
        {
            Program.RenderSize = renderSizeBox.Text;
            Console.WriteLine("RenderSize" + Program.RenderSize);
        }

        private void renderSizeBox_KeyPress(object sender, KeyPressEventArgs e)
        {
            if (!char.IsControl(e.KeyChar) && !char.IsDigit(e.KeyChar) &&
                (e.KeyChar != '.'))
            {
                e.Handled = true;
            }

            // only allow one decimal point
            if ((e.KeyChar == '.') && ((sender as TextBox).Text.IndexOf('.') > -1))
            {
                e.Handled = true;
            }
        }

        private void label2_Click(object sender, EventArgs e)
        {

        }

        private void button5_Click(object sender, EventArgs e)
        {
            FolderBrowserDialog folderBrowserDialog1 = new FolderBrowserDialog();
            if (folderBrowserDialog1.ShowDialog() == DialogResult.OK)
            {
                Program.Output = folderBrowserDialog1.SelectedPath;
            }
        }

        private void checkBox1_CheckedChanged(object sender, EventArgs e)
        {
            Program.ReadIDMap = checkBox1.Checked;
            Console.WriteLine(Program.ReadIDMap);
        }

        private void label2_Click_1(object sender, EventArgs e)
        {

        }

        private void maxFrontTextBox_TextChanged(object sender, EventArgs e)
        {
            Program.MaxFrontDistance = maxFrontTextBox.Text;
            Console.WriteLine("MaxFrontDistance" + Program.MaxFrontDistance);
        }

        private void maxFrontTextBox_KeyPress(object sender, KeyPressEventArgs e)
        {
            if (!char.IsControl(e.KeyChar) && !char.IsDigit(e.KeyChar) &&
                (e.KeyChar != '.'))
            {
                e.Handled = true;
            }

            // only allow five decimal point
            if ((e.KeyChar == '.') && ((sender as TextBox).Text.IndexOf('.') > -5))
            {
                e.Handled = true;
            }
        }

        private void maxRearTextBox_TextChanged(object sender, EventArgs e)
        {
            Program.MaxRearDistance = maxRearTextBox.Text;
            Console.WriteLine("MaxRearDistance" + Program.MaxRearDistance);
        }

        private void maxRearTextBox_KeyPress(object sender, KeyPressEventArgs e)
        {
            if (!char.IsControl(e.KeyChar) && !char.IsDigit(e.KeyChar) &&
                (e.KeyChar != '.'))
            {
                e.Handled = true;
            }

            // only allow five decimal point
            if ((e.KeyChar == '.') && ((sender as TextBox).Text.IndexOf('.') > -5))
            {
                e.Handled = true;
            }
        }
    }
}
