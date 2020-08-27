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
                Program.PLYFile = ofd.FileName;
            }
        }

        private void button4_Click(object sender, EventArgs e)
        {
            Console.WriteLine("BAKING");
            if (Program.LowPolyFile == "" || Program.HighPolyFile == "" || Program.PLYFile == "" || Program.Output == "" || Program.RenderSize == "")
            {
                MessageBox.Show("You did not set some arguments. Double check again",    "Baker error");
            }
            else
            {
                Process p = new Process();
                p.StartInfo.FileName = @"gobaker.exe";

                string arguments = "";
                arguments += " -l " + Program.LowPolyFile;
                arguments += " -h " + Program.HighPolyFile;
                arguments += " -hp " + Program.PLYFile;
                arguments += " -s " + Program.RenderSize;
                arguments += " -o " + Program.Output;

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

        private void textBox1_TextChanged(object sender, EventArgs e)
        {
            Program.RenderSize = textBox1.Text;
            Console.WriteLine(Program.RenderSize);
        }

        private void textBox1_KeyPress(object sender, KeyPressEventArgs e)
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
    }
}
