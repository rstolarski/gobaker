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
        public static string HighPolyPLYFile = "";
        public static string RenderSize = "";
        public static string MaxFrontDistance = "";
        public static string MaxRearDistance = "";
        public static string Output = "";
        public static bool ReadIDMap;
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
