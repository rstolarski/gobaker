using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace GoBaker_App
{
    static class Program
    {
        public static string LowPolyFile = "";
        public static string HighPolyFile = "";
        public static string PLYFile = "";
        public static string RenderSize = "";
        public static string Output = "";
        /// <summary>
        /// The main entry point for the application.
        /// </summary>
        [STAThread]
        static void Main()
        {
            Application.EnableVisualStyles();
            Application.SetCompatibleTextRenderingDefault(false);
            Application.Run(new Form1());
        }
    }
}
